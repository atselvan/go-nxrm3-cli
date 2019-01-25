import groovy.json.JsonOutput
import groovy.json.JsonSlurper
import org.sonatype.nexus.repository.maven.LayoutPolicy
import org.sonatype.nexus.repository.maven.VersionPolicy
import org.sonatype.nexus.repository.storage.WritePolicy

def input = new JsonSlurper().parseText(args)
def output = [:]

def isExists = repository.getRepositoryManager().get(input.name)

if (!isExists) {
    def repo = repository.createMavenHosted(input.name, "default", true, VersionPolicy.SNAPSHOT, WritePolicy.ALLOW, LayoutPolicy.STRICT)
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
