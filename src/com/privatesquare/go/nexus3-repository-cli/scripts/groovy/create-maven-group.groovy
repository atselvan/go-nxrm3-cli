import groovy.json.JsonOutput
import groovy.json.JsonSlurper
import org.sonatype.nexus.repository.maven.LayoutPolicy
import org.sonatype.nexus.repository.maven.VersionPolicy

def input = new JsonSlurper().parseText(args)
def output = [:]
def repo

if (!repository.getRepositoryManager().exists(input.name)){
    repo = repository.createMavenGroup(input.name, input.members, input.blobStoreName)
    output.put("status", "200 OK")
    output.put("name", repo.name)
    output.put("url", repo.url)
} else {
    output.put("status", "302 Found")
}

return JsonOutput.toJson(output)