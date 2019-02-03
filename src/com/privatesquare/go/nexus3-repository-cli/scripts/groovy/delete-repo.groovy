import groovy.json.JsonOutput
import groovy.json.JsonSlurper

def input = new JsonSlurper().parseText(args)

def output = [:]

if (repository.getRepositoryManager().exists(input.name)){
    output.put("status", "200 OK")
    repository.getRepositoryManager().delete(input.name)
    return JsonOutput.toJson(output)
} else {
    output.put("status", "404 Not Found")
    return JsonOutput.toJson(output)
}