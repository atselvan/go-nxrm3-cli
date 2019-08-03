// delete-role.groovy is a  Nexus3 Integration API definition to delete a repository role  from Nexus

// import libraries for json parsing
import groovy.json.JsonOutput
import groovy.json.JsonSlurper
// import Nexus functions from the nexus-security jar file
import org.sonatype.nexus.security.authz.AuthorizationManager

// input map
Map input = new JsonSlurper().parseText(args)
//  output map
Map output = [:]

authManager = container.lookup(AuthorizationManager)

authManager.deleteRole(input.roleId)

// output success status
output.put('status', '200 OK')

// nexus logger
log.info('***********************************************')
log.info(String.format('Role %s is deleted', input.roleId))
log.info('**********************************************')

// return  output in JSON format
JsonOutput.toJson(output)
