// delete-content-selector.groovy is a  Nexus3 Integration API definition to delete a content selector from Nexus

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

config = selectorManager.browse().find { it -> it.name == input.name }

selectorManager.delete(config)

// output success status
output.put('status', '200 OK')

// nexus  logger
log.info('***********************************************')
log.info(String.format('Content selector %s is deleted', input.name))
log.info('***********************************************')

// return output in JSON format
return JsonOutput.toJson(output)
