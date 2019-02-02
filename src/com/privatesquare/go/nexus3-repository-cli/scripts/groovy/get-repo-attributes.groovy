import groovy.json.JsonOutput
import groovy.json.JsonSlurper

def input = new JsonSlurper().parseText(args)
def output = [:]
def repo
def configuration
def attributes

//if (repository.getRepositoryManager().exists(input.name)){
    //repo = repository.createMavenProxy(input.name, input.remoteURL, input.blobStoreName, false, VersionPolicy.RELEASE, LayoutPolicy.PERMISSIVE)

    log.info("***********************************************")
    log.info("Getting repository attributes")
    log.info("***********************************************")

    //repo = repository.getRepositoryManager().get(input.name)
    //configuration = repo.getConfiguration()
    //attributes = configuration.getAttributes()

    output.put("status", "200 OK")
    //output.put("name", configuration.getRepositoryName())

    //attributes.remove("something")
    //output.put("test", attributes)
    //output.put("test1", "something")


    output.put("name", input.attribute.maven.versionPolicy)

    return JsonOutput.toJson(output)

//} else {
//    output.put("status", "404 Not Found")
//    return JsonOutput.toJson(output)
//}


