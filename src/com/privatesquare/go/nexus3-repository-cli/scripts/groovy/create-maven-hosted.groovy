import groovy.json.JsonOutput
import groovy.json.JsonSlurper
import org.sonatype.nexus.repository.storage.WritePolicy

def output = [:]
def input = new JsonSlurper().parseText(args)

def isExists = repository.getRepositoryManager().get(input.name)

if (!isExists) {
    def repo = repository.createMavenHosted(input.name, "default", true, VersionPolicy.SNAPSHOTS, WritePolicy.ALLOW, LayoutPolicy.STRICT)
    output.put("status", "200")
    output.put("name", repo.name)
    output.put("url", repo.url)
    output.put("type", repo.type.value)
    output.put("format", repo.format.value)
    output.put("recipe", repo.configuration.recipeName)
} else {
    output.put("status", "409")
    output.put("message", "Repository already exists!")
}

return JsonOutput.toJson(output)
