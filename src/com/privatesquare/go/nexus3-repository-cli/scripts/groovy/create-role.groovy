// create-role.groovy is a  Nexus3 Integration API definition to create a repository role  in Nexus

// import libraries for json parsing
import groovy.json.JsonOutput
import groovy.json.JsonSlurper
// import Nexus libraries from the nexus-security jar file
import org.sonatype.nexus.security.authz.AuthorizationManager
import org.sonatype.nexus.security.role.Role

// input map
Map input = new JsonSlurper().parseText(args)
// output map
Map output = [:]

authManager = container.lookup(AuthorizationManager)

role = new Role(
        roleId: input.roleId,
        name: input.roleId,
        description: input.description,
        source: input.source,
        roles: input.roles,
        privileges: input.privileges,
)

authManager.addRole(role)

// output success status
output.put('status', '200 OK')
output.put('role', role)

// nexus logger
log.info('***********************************************')
log.info(String.format('Role %s is created', input.name))
log.info('**********************************************')

// return output in JSON format
JsonOutput.toJson(output)
