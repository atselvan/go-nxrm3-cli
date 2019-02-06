package backend

import (
	m "com/privatesquare/go/nexus3-repository-cli/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ListSelectors(name string) {
	if name != "" {
		cs := getSelector(name)
		fmt.Printf("Name: %s\nDescription: %s\nExpression: %s\n",
			cs.Name, cs.Description, cs.Attributes.Expression)
	} else {
		csNames := getSelectorNames()
		printStringSlice(csNames)
		fmt.Printf("Total Number of content selectors : %d\n", len(csNames))
	}
}

func CreateSelector(name, description, expression string) {
	if name == "" || expression == "" {
		log.Printf("%s : %s", getfuncName(), createSelectorRequiredInfo)
		os.Exit(1)
	}
	if !selectorExists(name) {
		attributes := m.ContentSelectorAttributes{Expression: expression}
		payload, err := json.Marshal(m.ContentSelector{Name: name, Type: contentSelectorType, Description: getSelectorDescription(description), Attributes: attributes})
		logJsonMarshalError(err, jsonMarshalError)
		result := RunScript("create-content-selector", string(payload))
		if result.Status == "200 OK" {
			log.Printf(createSelectorSuccessInfo, name)
		} else {
			log.Printf("%s : %s", getfuncName(), setVerboseInfo)
			os.Exit(1)
		}
	} else {
		log.Printf(selectorAlreadyExistsInfo, name)
	}
}

// TODO : Fix the java.lang.IllegalStateException: Missing entity-metadata during runtime
func UpdateSelector(name, description, expression string) {
	if name == "" {
		log.Printf("%s : %s", getfuncName(), selectorNameRequiredInfo)
		os.Exit(1)
	}
	if selectorExists(name) {
		selector := getSelector(name)
		if description != "" {
			selector.Description = description
		}
		if expression != "" {
			selector.Attributes = m.ContentSelectorAttributes{Expression: expression}
		}
		payload, err := json.Marshal(selector)
		logJsonMarshalError(err, jsonMarshalError)
		result := RunScript("update-content-selector", string(payload))
		if result.Status == "200 OK" {
			log.Printf(updateSelectorSuccessInfo, name)
		} else {
			log.Printf("%s : %s", getfuncName(), setVerboseInfo)
			os.Exit(1)
		}
	} else {
		log.Printf(selectorNotFoundInfo, name)
	}
}

// TODO : Fix the java.lang.IllegalStateException: Missing entity-metadata during runtime
func DeleteSelector(name string) {
	if name == "" {
		log.Printf("%s : %s", getfuncName(), selectorNameRequiredInfo)
		os.Exit(1)
	}
	if selectorExists(name) {
		selector := getSelector(name)
		payload, err := json.Marshal(selector)
		logJsonMarshalError(err, jsonMarshalError)
		result := RunScript("delete-content-selector", string(payload))
		if result.Status == "200 OK" {
			log.Printf(deleteSelectorSuccessInfo, name)
		} else {
			log.Printf("%s : %s", getfuncName(), setVerboseInfo)
			os.Exit(1)
		}
	} else {
		log.Printf(selectorNotFoundInfo, name)
	}
}

func getSelectors() []m.ContentSelector {
	payload, err := json.Marshal(m.Repository{})
	logJsonMarshalError(err, getfuncName())
	result := RunScript("get-content-selectors", string(payload))
	return result.ContentSelectors
}

func getSelector(name string) m.ContentSelector {
	if name == "" {
		log.Printf("%s : %s", getfuncName(), selectorNameRequiredInfo)
		os.Exit(1)
	}
	var contentSelector m.ContentSelector
	contentSelectors := getSelectors()
	for _, cs := range contentSelectors {
		if cs.Name == name {
			contentSelector = cs
		}
	}
	if contentSelector.Name == "" {
		log.Printf(selectorNotFoundInfo, name)
		os.Exit(1)
	}
	return contentSelector
}

func getSelectorNames() []string {
	var csNames []string
	contentSelectors := getSelectors()
	for _, cs := range contentSelectors {
		csNames = append(csNames, cs.Name)
	}
	return csNames
}

func selectorExists(name string) bool {
	csNames := getSelectorNames()
	if entryExists(csNames, name) {
		return true
	}
	return false
}

func getSelectorDescription(description string) string {
	if description == "" {
		return "custom content selector"
	}
	return description
}
