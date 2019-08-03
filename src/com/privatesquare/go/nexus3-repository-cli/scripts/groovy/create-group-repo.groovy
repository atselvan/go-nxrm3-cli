// create-group-repo.groovy is a  Nexus3 Integration API definition to create a group repository in Nexus

// import libraries for json parsing
import groovy.json.JsonOutput
import groovy.json.JsonSlurper
// import Nexus Configuration function from nexus-repository jar file
import org.sonatype.nexus.repository.config.Configuration

// input map
Map input = new JsonSlurper().parseText(args)
// output map
Map output = [:]
Configuration configuration

// reference : https://help.sonatype.com/repomanager3/configuration/repository-management

// maven configuration
// reference : https://help.sonatype.com/repomanager3/maven-repositories
Map mavenConfig = [
        'versionPolicy': input.attributes.maven.versionPolicy,
        'layoutPolicy': input.attributes.maven.layoutPolicy,
]

// storage configuration
Map storageConfig = [
        'blobStoreName': input.attributes.storage.blobStoreName,
        'strictContentTypeValidation': input.attributes.storage.strictContentTypeValidation,
        'writePolicy': input.attributes.storage.writePolicy
]

// clean up policy configuration
Map cleanUpConfig = [
        'policyName': 'None'
]

// docker configuration
Map dockerConfig = [
        'forceBasicAuth': input.attributes.docker.forceBasicAuth,
        'v1Enabled': input.attributes.docker.v1Enabled
]
if (input.attributes.docker.httpPort == 0){
    dockerConfig.put('httpsPort', input.attributes.docker.httpsPort)
} else if (input.attributes.docker.httpsPort == 0){
    dockerConfig.put('httpPort', input.attributes.docker.httpPort)
} else {
    dockerConfig.put('httpPort', input.attributes.docker.httpPort)
    dockerConfig.put('httpsPort', input.attributes.docker.httpsPort)
}

// group repo members config
Map groupConfig = [
        'memberNames': input.attributes.group.memberNames
]

// check  if repository does not exists before creating a repository
if (!repository.getRepositoryManager().exists(input.name)){
    configuration = new Configuration()
    if (input.format == 'maven2'){
        configuration.setAttributes(
                'maven': mavenConfig,
                'storage': storageConfig,
                'group': groupConfig,
                'cleanup': cleanUpConfig
        )
    } else if (input.format == 'docker'){
        configuration.setAttributes(
                'docker': dockerConfig,
                'storage': storageConfig,
                'group': groupConfig,
                'cleanup': cleanUpConfig
        )
    } else {
        configuration.setAttributes(
                'storage': storageConfig,
                'group': groupConfig,
                'cleanup': cleanUpConfig
        )
    }
    configuration.setRepositoryName(input.name)
    configuration.setRecipeName(input.recipe)
    configuration.setOnline(true)
    repo = repository.repositoryManager.create(configuration)
    attributes = repo.getConfiguration().getAttributes()

    // output success request status
    output.put('status', '200 OK')
    output.put('name', repo.name)
    output.put('url', repo.url)
    output.put('recipe', repo.configuration.recipeName)
    output.put('attributes', attributes)

    // nexus logger
    log.info('**********************************************')
    log.info(String.format('Repository %s is created!!!', repo.name))
    log.info('*********************************************')

    return JsonOutput.toJson(output)
} else {
    // output found request status
    output.put('status', '302 Found')
    return JsonOutput.toJson(output)
}
