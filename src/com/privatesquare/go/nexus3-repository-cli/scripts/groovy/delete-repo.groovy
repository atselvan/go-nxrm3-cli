// delete-repo.groovy is a  Nexus3 Integration API definition to delete a repository from Nexus

// import libraries for json parsing
import groovy.json.JsonOutput
import groovy.json.JsonSlurper

// input map
Map input = new JsonSlurper().parseText(args)

// output map
Map output = [:]

if (repository.getRepositoryManager().exists(input.name)){
    // output success status
    output.put("status", "200 OK")
    repository.getRepositoryManager().delete(input.name)
    // return output in JSON format
    return JsonOutput.toJson(output)
} else {
    // output not found status
    output.put("status", "404 Not Found")
    // return output in JSON format
    return JsonOutput.toJson(output)
}
