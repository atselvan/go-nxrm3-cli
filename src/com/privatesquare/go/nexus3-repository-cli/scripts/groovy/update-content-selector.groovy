// update-content-selector.groovy is a  Nexus3 Integration API definition to update a content selector in Nexus

// import libraries for json parsing
import groovy.json.JsonOutput
import groovy.json.JsonSlurper
// import Nexus SelectorManager function from nexus-selector log file
import org.sonatype.nexus.selector.SelectorManager

// input map
Map input = new JsonSlurper().parseText(args)
// output map
Map output = [:]

selectorManager = container.lookup(SelectorManager)

configuration = selectorManager.browse().find { it -> it.name == input.name }

configuration.setDescription(input.description)

configuration.setAttributes(input.attributes)

selectorManager.update(configuration)

// output success status
output.put('status', '200 OK')

//nexus logger
log.info('***********************************************')
log.info(String.format('Content selector %s is updated', input.name))
log.info('**********************************************')

// return output in JSON format
JsonOutput.toJson(output)
