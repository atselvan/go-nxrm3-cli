import groovy.json.JsonOutput
import groovy.json.JsonSlurper
import org.sonatype.nexus.repository.*
import org.sonatype.nexus.repository.maven.LayoutPolicy
import org.sonatype.nexus.repository.maven.VersionPolicy

log.info("***********************************************")
log.info("***********************************************")

def input = new JsonSlurper().parseText(args)
def output = [:]
def configuration
def repo

def username = "test"
def password = "test123"

if (!repository.getRepositoryManager().exists(input.name)){
    configuration = new org.sonatype.nexus.repository.config.Configuration()
    configuration.setAttributes(
            'maven': [
                    'versionPolicy': 'RELEASE',
                    'layoutPolicy': 'STRICT'
            ],
            'proxy': [
                    'remoteUrl': input.remoteURL,
                    'contentMaxAge': -1,
                    'metadataMaxAge': 1440
            ],
            'httpclient': [
                    'blocked': false,
                    'autoBlock': true,
                    'authentication': [
                            'type': 'username',
                            'username': username,
                            'password': password
                    ]
            ],
            'storage': [
                    'blobStoreName': 'default',
                    'strictContentTypeValidation': true
            ]
    )
    configuration.setRepositoryName(input.name)
    configuration.setRecipeName('maven2-proxy')
    configuration.setOnline(true)

    repo = repository.repositoryManager.create(configuration)
    attributes = repo.getConfiguration().getAttributes()

    output.put("status", (Object)"200 OK")
    output.put("name", (Object)repo.name)
    output.put("url", (Object)repo.url)
    log.info("***********************************************")
    log.info("Repository Created!!!")
    log.info("***********************************************")
} else {
    output.put("status", "302 Found")
}

return JsonOutput.toJson(attributes)