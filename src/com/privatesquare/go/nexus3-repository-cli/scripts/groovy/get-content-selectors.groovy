import groovy.json.JsonOutput
import org.sonatype.nexus.selector.*

def contentSlectorList = []
output = [:]

def selectorManager = container.lookup(SelectorManager.class.name)

selectorManager.browse().each{ cs ->
    def contentSelector = [:]
    contentSelector.put("name", cs.name)
    contentSelector.put("type", cs.type)
    contentSelector.put("description", cs.description)
    contentSelector.put("attributes", cs.attributes)
    contentSlectorList.add(contentSelector)
}

output.put("status", "200 OK")
output.put("contentSelectors", contentSlectorList)

log.info("***********************************************")
log.info(String.format("Testing security classes"))
log.info("***********************************************")

return JsonOutput.toJson(output)