import groovy.json.JsonOutput
import org.sonatype.nexus.security.authz.*

output = [:]

authManager = container.lookup(AuthorizationManager.class.name)
roles = authManager.listRoles()

output.put("status", "200 OK")
output.put("roles", roles)

return JsonOutput.toJson(output)
