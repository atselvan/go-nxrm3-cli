import groovy.json.JsonOutput
import groovy.json.JsonSlurper

def input = new JsonSlurper().parseText(args)
def repo = repository.getRepositoryManager().get(input.name)

def output = [:]

if (repository.getRepositoryManager().exists(input.name)){
    output.put("status", "200 OK")
    output.put("name", repo.name)
    output.put("url", repo.url)
    output.put("type", repo.type.value)
    output.put("format", repo.format.value)
    output.put("recipe", repo.configuration.recipeName)
    output.put("attributes", repo.configuration.attributes)
} else {
    output.put("status", "404 Not Found")
}

return JsonOutput.toJson(output)