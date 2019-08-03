// get-repo.groovy is a  Nexus3 Integration API definition to get repositories from Nexus

// import libraries for json parsing
import groovy.json.JsonOutput
import groovy.json.JsonSlurper

// input map
Map input = new JsonSlurper().parseText(args)

// getting repositories from nexus
repo = repository.getRepositoryManager().get(input.name)

// output map
Map output = [:]

if (repository.getRepositoryManager().exists(input.name)){
    // return success status
    output.put("status", "200 OK")
    output.put("name", repo.name)
    output.put("url", repo.url)
    output.put("type", repo.type.value)
    output.put("format", repo.format.value)
    output.put("recipe", repo.configuration.recipeName)
    output.put("attributes", repo.configuration.attributes)
} else {
    // return not found status
    output.put("status", "404 Not Found")
}

// return output in JSON format
JsonOutput.toJson(output)
