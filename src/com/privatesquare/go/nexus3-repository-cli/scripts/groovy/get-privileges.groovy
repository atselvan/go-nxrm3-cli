// get-privileges.groovy is a  Nexus3 Integration API definition to get all repository privileges from Nexus

// import libraries for json parsing
import groovy.json.JsonOutput
// import nexus libraries from the nexus-security jar file
import org.sonatype.nexus.security.authz.AuthorizationManager

// output map
Map output = [:]

authManager = container.lookup(AuthorizationManager)
privileges = authManager.listPrivileges()

// output success status
output.put('status', '200 OK')
output.put('privileges', privileges)

// return output in JSON format
JsonOutput.toJson(output)
