import groovy.json.JsonOutput
import groovy.json.JsonSlurper
import org.sonatype.nexus.security.authz.*

def input = new JsonSlurper().parseText(args)
output = [:]

def authManager = container.lookup(AuthorizationManager.class.name)

authManager.deleteRole(input.roleId)
output.put("status", "200 OK")

log.info("***********************************************")
log.info(String.format("Role %s is deleted", input.roleId))
log.info("***********************************************")

return JsonOutput.toJson(output)


