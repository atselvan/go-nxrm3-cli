import groovy.json.JsonOutput
import groovy.json.JsonSlurper
import org.sonatype.nexus.security.authz.*

def input = new JsonSlurper().parseText(args)
output = [:]

def authManager = container.lookup(AuthorizationManager.class.name)

authManager.deletePrivilege(input.id)
output.put("status", "200 OK")

log.info("***********************************************")
log.info(String.format("Privilege %s is deleted", input.id))
log.info("***********************************************")

return JsonOutput.toJson(output)

