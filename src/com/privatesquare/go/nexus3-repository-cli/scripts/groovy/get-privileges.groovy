import groovy.json.JsonOutput
import org.sonatype.nexus.security.authz.*

output = [:]

authManager = container.lookup(AuthorizationManager.class.name)
privileges = authManager.listPrivileges()

output.put("status", "200 OK")
output.put("privileges", privileges)

return JsonOutput.toJson(output)