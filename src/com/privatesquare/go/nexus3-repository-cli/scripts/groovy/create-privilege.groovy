import groovy.json.JsonOutput
import groovy.json.JsonSlurper
import org.sonatype.nexus.security.authz.*
import org.sonatype.nexus.security.privilege.*

def input = new JsonSlurper().parseText(args)
output = [:]

def authManager = container.lookup(AuthorizationManager.class.name)

def properties = [:]
properties.put("contentSelector", input.properties.contentSelector)
properties.put("repository", input.properties.repository)
properties.put("actions", input.properties.actions)

privilege = new Privilege(
        id: input.id,
        name: input.name,
        description: input.description,
        type: input.type,
        version: '',
        properties: properties
)

authManager.addPrivilege(privilege)

output.put("status", "200 OK")

log.info("***********************************************")
log.info(String.format("Privilege %s is created", input.name))
log.info("***********************************************")

return JsonOutput.toJson(output)
