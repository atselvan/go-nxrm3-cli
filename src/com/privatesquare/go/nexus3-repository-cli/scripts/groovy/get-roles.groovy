// get-roles.groovy is a  Nexus3 Integration API definition to get all the roles that are available in Nexus

// import libraries for json parsing
import groovy.json.JsonOutput
// Importing Nexus AuthorizationManager function from nexus-security jar file
import org.sonatype.nexus.security.authz.AuthorizationManager

// output map for return values
Map output = [:]

authManager = container.lookup(AuthorizationManager)
roles = authManager.listRoles()

// output the status of the request
output.put('status', '200 OK')
// output the roles from nexus
output.put('roles', roles)

// return output in JSON format
JsonOutput.toJson(output)
