import groovy.json.JsonOutput
import groovy.json.JsonSlurper
import org.sonatype.nexus.repository.maven.LayoutPolicy
import org.sonatype.nexus.repository.maven.VersionPolicy
import org.sonatype.nexus.repository.storage.WritePolicy

def input = new JsonSlurper().parseText(args)
def output = [:]
def repo

if (!repository.getRepositoryManager().exists(input.name)){
    if (input.versionPolicy == "release"){
        repo = repository.createMavenHosted(input.name, input.blobStoreName, true, VersionPolicy.RELEASE, WritePolicy.ALLOW_ONCE, LayoutPolicy.STRICT)
    } else {
        repo = repository.createMavenHosted(input.name, input.blobStoreName, true, VersionPolicy.SNAPSHOT, WritePolicy.ALLOW, LayoutPolicy.STRICT)
    }
    output.put("status", "200 OK")
    output.put("name", repo.name)
    output.put("url", repo.url)
} else {
    output.put("status", "302 Found")
}

return JsonOutput.toJson(output)