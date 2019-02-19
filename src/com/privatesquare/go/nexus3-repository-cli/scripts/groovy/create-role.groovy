import groovy.json.JsonOutput
import groovy.json.JsonSlurper
import org.sonatype.nexus.security.authz.*
import org.sonatype.nexus.security.role.*

def input = new JsonSlurper().parseText(args)
output = [:]

def authManager = container.lookup(AuthorizationManager.class.name)

role = new Role(
        roleId: input.roleId,
        name: input.roleId,
        description: input.description,
        source: input.source,
        roles: input.roles,
        privileges: input.privileges,
)

authManager.addRole(role)

output.put("status", "200 OK")
output.put("role", role)

log.info("***********************************************")
log.info(String.format("Role %s is created", input.name))
log.info("***********************************************")

return JsonOutput.toJson(output)
