import groovy.json.JsonOutput
import groovy.json.JsonSlurper

log.info("***********************************************")
log.info("Executing create-hosted-repo script")
log.info("***********************************************")

def input = new JsonSlurper().parseText(args)
def output = [:]
def configuration
def repo

def mavenConfig = [
        'versionPolicy': input.attributes.maven.versionPolicy,
        'layoutPolicy': input.attributes.maven.layoutPolicy,
]

def storageConfig = [
        'blobStoreName': input.attributes.storage.blobStoreName,
        'strictContentTypeValidation': input.attributes.storage.strictContentTypeValidation,
        'writePolicy': input.attributes.storage.writePolicy
]

def cleanUpConfig = [
        'policyName': 'None'
]

def dockerConfig = [
        "forceBasicAuth": input.attributes.docker.forceBasicAuth,
        "v1Enabled": input.attributes.docker.v1Enabled
]
if (input.attributes.docker.httpPort == 0){
    dockerConfig.put("httpsPort", input.attributes.docker.httpsPort)
} else if (input.attributes.docker.httpsPort == 0){
    dockerConfig.put("httpPort", input.attributes.docker.httpPort)
} else {
    dockerConfig.put("httpPort", input.attributes.docker.httpPort)
    dockerConfig.put("httpsPort", input.attributes.docker.httpsPort)
}

if (!repository.getRepositoryManager().exists(input.name)){
    configuration = new org.sonatype.nexus.repository.config.Configuration()
    if (input.format == "maven2"){
        configuration.setAttributes(
                'maven': mavenConfig,
                'storage': storageConfig,
                'cleanup': cleanUpConfig
        )
    } else if (input.format == "docker"){
        configuration.setAttributes(
                'docker': dockerConfig,
                'storage': storageConfig,
                'cleanup': cleanUpConfig
        )
    } else {
        configuration.setAttributes(
                'storage': storageConfig,
                'cleanup': cleanUpConfig
        )
    }
    configuration.setRepositoryName(input.name)
    configuration.setRecipeName(input.recipe)
    configuration.setOnline(true)
    repo = repository.repositoryManager.create(configuration)
    attributes = repo.getConfiguration().getAttributes()

    output.put("status", "200 OK")
    output.put("name", repo.name)
    output.put("url", repo.url)
    output.put("recipe", repo.configuration.recipeName)
    output.put("attributes", attributes)

    log.info("***********************************************")
    log.info(String.format("Repository %s is created!!!", repo.name))
    log.info("***********************************************")

    return JsonOutput.toJson(output)
} else {
    output.put("status", "302 Found")
    return JsonOutput.toJson(output)
}


