import groovy.json.JsonOutput
import groovy.json.JsonSlurper
import org.sonatype.nexus.selector.*

def input = new JsonSlurper().parseText(args)
output = [:]

def selectorManager = container.lookup(SelectorManager.class.name)

configuration = new SelectorConfiguration(
        name: input.name,
        type: input.type,
        description: input.description,
        attributes: input.attributes
)

selectorManager.create(configuration)

output.put("status", "200 OK")

log.info("***********************************************")
log.info(String.format("Content selector %s is created", input.name))
log.info("***********************************************")

return JsonOutput.toJson(output)