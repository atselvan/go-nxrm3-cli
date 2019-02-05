import groovy.json.JsonOutput
import groovy.json.JsonSlurper
import org.sonatype.nexus.selector.*

def input = new JsonSlurper().parseText(args)
output = [:]

def selectorManager = container.lookup(SelectorManager.class.name)

configuration = new SelectorConfiguration(
        name: input.name,
        type: "csel",
        description: input.description,
        attributes: input.attributes
)

selectorManager.update(configuration)

output.put("status", "200 OK")

log.info("***********************************************")
log.info(String.format("Content selector %s is updated", input.name))
log.info("***********************************************")

return JsonOutput.toJson(output)
