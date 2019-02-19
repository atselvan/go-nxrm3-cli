package backend

import (
	m "com/privatesquare/go/nexus3-repository-cli/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ListPrivileges(name string) {
	if name != "" {
		privilege := getPrivilege(name)
		fmt.Printf("%+v\n", privilege)
	} else {
		pNames := getPrivilegeNames()
		printStringSlice(pNames)
		fmt.Printf("Number of privileges in nexus : %d\n", len(pNames))
	}
}

func CreatePrivilege(name, description, selectorName, repoName, action string) {
	if name == "" || selectorName == "" || repoName == "" {
		log.Printf("%s : %s", getfuncName(), createPrivilegeRequiredInfo)
		os.Exit(1)
	}
	if !privilegeExists(name) {
		properties := m.PrivilegeProperties{ContentSelector: validateSelectorForPriv(selectorName), Repository: validateRepoForPriv(repoName), Actions: getPrivilegeActions(action)}
		payload, err := json.Marshal(m.Privilege{ID: toLower(name), Name: toLower(name), Description: getPrivilegeDescription(description), Type: getPrivilegeType(), Properties: properties, ReadOnly: false})
		logJsonMarshalError(err, jsonMarshalError)
		result := RunScript(createPrivilegeScript, string(payload))
		if result.Status == successStatus {
			log.Printf(createPrivilegeSuccessInfo, name)
		}
	} else {
		log.Printf(privilegeExistsInfo, name)
	}
}

func UpdatePrivilege(name, description, selectorName, repoName, action string) {
	if name == "" {
		log.Printf("%s : %s", getfuncName(), privilegeNameRequiredInfo)
		os.Exit(1)
	}
	privilege := getPrivilege(name)
	if description != "" {
		privilege.Description = description
	}
	if selectorName != "" {
		privilege.Properties.ContentSelector = validateSelectorForPriv(selectorName)
	}
	if repoName != "" {
		privilege.Properties.Repository = validateRepoForPriv(repoName)
	}
	if action != "" {
		privilege.Properties.Actions = getPrivilegeActions(action)
	}
	if privilegeExists(name) {
		payload, err := json.Marshal(privilege)
		logJsonMarshalError(err, jsonMarshalError)
		result := RunScript(updatePrivilegeScript, string(payload))
		if result.Status == successStatus {
			log.Printf(updatePrivilegeSuccessInfo, name)
		}
	} else {
		log.Printf(privilegeNotFoundInfo, name)
	}
}

func DeletePrivilege(name string) {
	if name == "" {
		log.Printf("%s : %s", getfuncName(), privilegeNameRequiredInfo)
		os.Exit(1)
	}
	if privilegeExists(name) {
		payload, err := json.Marshal(m.Privilege{ID: getPrivilegeID(name)})
		logJsonMarshalError(err, jsonMarshalError)
		result := RunScript(deletePrivilegeScript, string(payload))
		if result.Status == successStatus {
			log.Printf(deletePrivilegeSuccessInfo, name)
		}
	} else {
		log.Printf(privilegeNotFoundInfo, name)
	}
}

func getPrivileges() []m.Privilege {
	payload, err := json.Marshal(m.Privilege{})
	logJsonMarshalError(err, getfuncName())
	result := RunScript(getPrivilegesScript, string(payload))
	return result.Privileges
}

func getPrivilege(name string) m.Privilege {
	if name == "" {
		log.Printf("%s : %s", getfuncName(), privilegeNameRequiredInfo)
		os.Exit(1)
	}
	var privilege m.Privilege
	privileges := getPrivileges()
	for _, p := range privileges {
		if p.Name == name {
			privilege = p
		}
	}
	if privilege.Name == "" {
		log.Printf(privilegeNotFoundInfo, name)
		os.Exit(1)
	}
	return privilege
}

func getPrivilegeNames() []string {
	var pNames []string
	privileges := getPrivileges()
	for _, p := range privileges {
		pNames = append(pNames, p.Name)
	}
	return pNames
}

func getPrivilegeID(name string) string {
	return getPrivilege(name).ID
}

func privilegeExists(name string) bool {
	pNames := getPrivilegeNames()
	if entryExists(pNames, name) {
		return true
	}
	return false
}

func getPrivilegeType() string {
	return "repository-content-selector"
}

func getPrivilegeDescription(description string) string {
	if description == "" {
		return defaultPrivilegeDescription
	}
	return description
}

func getPrivilegeActions(action string) string {
	if action == "read" {
		return "browse,read"
	} else if action == "write" {
		return "add,browse,create,edit,read,update"
	} else {
		return "*"
	}
}

func validateSelectorForPriv(selectorName string) string {
	if !selectorExists(selectorName) {
		log.Printf("%s : "+selectorNotFoundInfo, getfuncName(), selectorName)
		os.Exit(1)
	}
	return selectorName
}

func validateRepoForPriv(repoName string) string {
	var allowedFormats []string
	allowedFormats = append(allowedFormats, "*")
	for _, format := range RepoFormats {
		allowedFormats = append(allowedFormats, fmt.Sprintf("*-%s", validateRepositoryFormat(format)))
	}
	if !entryExists(allowedFormats, repoName) {
		if !repositoryExists(repoName) {
			log.Printf("%s : "+repositoryNotFoundInfo, getfuncName(), repoName)
			os.Exit(1)
		}
	}
	return repoName
}
