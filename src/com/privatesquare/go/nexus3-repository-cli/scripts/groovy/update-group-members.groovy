import groovy.json.JsonOutput
import groovy.json.JsonSlurper

def input = new JsonSlurper().parseText(args)
def output = [:]
def configuration
def repo

if (repository.getRepositoryManager().exists(input.name)){
    repo = repository.getRepositoryManager().get(input.name)
    configuration = repo.getConfiguration()
    attributes = configuration.getAttributes()
    attributes.put('group', [
            'memberNames': input.attributes.group.memberNames
    ])

    repo = repository.repositoryManager.update(configuration)
    attributes = repo.getConfiguration().getAttributes()

    // output success status
    output.put("status", "200 OK")
    output.put("name", repo.name)
    output.put("url", repo.url)
    output.put("recipe", repo.configuration.recipeName)
    output.put("attributes", attributes)

    // nexus logger
    log.info("***********************************************")
    log.info(String.format("Repository %s is updated!!!", repo.name))
    log.info("**********************************************")
    // return output in JSON format
    return JsonOutput.toJson(output)
} else {
    // output not found status
    output.put("status", "404 Not Found")
    // return output in JSON format
    return JsonOutput.toJson(output)
}
