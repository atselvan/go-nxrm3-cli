// delete-privilege.groovy is a  Nexus3 Integration API definition to delete a repository privilege from Nexus

// import libraries for json parsing
import groovy.json.JsonOutput
import groovy.json.JsonSlurper
// import nexus libraries from the nexus-security jar file
import org.sonatype.nexus.security.authz.AuthorizationManager

// imput map
Map input = new JsonSlurper().parseText(args)
// output map
Map output = [:]

authManager = container.lookup(AuthorizationManager)

authManager.deletePrivilege(input.id)

// output success status
output.put('status', '200 OK')

// nexus logger
log.info('***********************************************')
log.info(String.format('Privilege %s is deleted', input.id))
log.info('**********************************************')

// return output in JSON format
JsonOutput.toJson(output)
