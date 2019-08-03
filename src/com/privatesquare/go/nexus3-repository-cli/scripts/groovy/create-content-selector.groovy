// create-content-selector.groovy is a  Nexus3 Integration API definition to create a content selector in Nexus

// import libraries for json parsing
import groovy.json.JsonOutput
import groovy.json.JsonSlurper
// Importing Nexus SelectorManager function from nexus-selector jar file
import org.sonatype.nexus.selector.SelectorManager
// Importing Nexus SelectorConfiguration function from nexus-selector jar file
import org.sonatype.nexus.selector.SelectorConfiguration

// input map
Map input = new JsonSlurper().parseText(args)

// output map
Map output = [:]

selectorManager = container.lookup(SelectorManager)

// selector configuration map
configMap = [
        name: input.name,
        type: input.type,
        description: input.description,
        attributes: input.attributes,
]

configuration = new SelectorConfiguration(configMap)

selectorManager.create(configuration)

// output request status
output.put('status', '200 OK')

// nexus logger
log.info('**********************************************')
log.info(String.format('Content selector %s is created', input.name))
log.info('********************************************')

// return output in JSON format
JsonOutput.toJson(output)
