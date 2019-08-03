// get-content-selectors.groovy is a  Nexus3 Integration API definition to get all content selectors from Nexus

// import libraries for json parsing
import groovy.json.JsonOutput
// import Nexus SelectorManager function from nexus-selector log file
import org.sonatype.nexus.selector.SelectorManager

List contentSlectorList = []
// output map
Map output = [:]

selectorManager = container.lookup(SelectorManager)

selectorManager.browse().each{ cs ->
    Map contentSelector = [:]
    contentSelector.put('name', cs.name)
    contentSelector.put('type', cs.type)
    contentSelector.put('description', cs.description)
    contentSelector.put('attributes', cs.attributes)
    contentSlectorList.add(contentSelector)
}

// output success status
output.put('status', '200 OK')
output.put('contentSelectors', contentSlectorList)

// return output in JSON format
JsonOutput.toJson(output)
