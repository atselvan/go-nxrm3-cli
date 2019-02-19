package backend

import (
	m "com/privatesquare/go/nexus3-repository-cli/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func ListRoles(id string) {
	if id != "" {
		role := getRole(id)
		fmt.Printf("Role Details:\n"+
			"ID: %s\n"+
			"Name: %s\n"+
			"Description: %s\n"+
			"Source: %s\n"+
			"Roles: %s\n"+
			"Privileges: %s\n",
			role.RoleID, role.Name, role.Description, role.Source, role.Roles, role.Privileges)
	} else {
		rIds := getRoleIDs()
		printStringSlice(rIds)
		log.Printf("Number of roles in nexus : %d\n", len(rIds))
	}
}

func CreateRole(id, description, roleMembers, rolePrivileges string) {
	if id == "" {
		log.Printf("%s : %s", getfuncName(), createRoleRequiredInfo)
		os.Exit(1)
	}
	validRoleMembers := validateRoleMembers(id, roleMembers)
	validRolePrivileges := validateRolePrivileges(rolePrivileges)
	if len(validRoleMembers)+len(validRolePrivileges) < 1 {
		log.Printf("%s : You need to provide atleast one valid role member or role privilege during role creation", getfuncName())
		os.Exit(1)
	}
	if !roleExists(id) {
		role := m.Role{RoleID: id, Name: id, Description: getRoleDesc(description), Source: getRoleSource(), Roles: validRoleMembers, Privileges: validRolePrivileges}
		payload, err := json.Marshal(role)
		logJsonMarshalError(err, jsonMarshalError)
		result := RunScript(createRoleScript, string(payload))
		if result.Status == successStatus {
			log.Printf(createRoleSuccessInfo, id, validRoleMembers, validRolePrivileges)
		} else {
			log.Printf("%s : %s", getfuncName(), setVerboseInfo)
		}
	} else {
		log.Printf(roleExistsInfo, id)
	}
}

func UpdateRole(id, description, roleMembers, rolePrivileges, updateAction string) {
	if id == "" {
		log.Printf("%s : %s", getfuncName(), roleIDRequiredInfo)
		os.Exit(1)
	}

	if updateAction == "" {
		log.Printf("%s : %s", getfuncName(), fmt.Sprintf(UpdateActionRequiredInfo, UpdateActions))
		os.Exit(1)
	}

	role := getRole(id)

	if description != "" {
		role.Description = description
	}

	validRoleMembers := validateRoleMembers(id, roleMembers)
	validRolePrivileges := validateRolePrivileges(rolePrivileges)

	if len(validRoleMembers)+len(validRolePrivileges) < 1 {
		log.Printf(roleItemsRequiredInfo, getfuncName())
		os.Exit(1)
	}

	if updateAction == "add" {
		for _, rm := range validRoleMembers {
			if !entryExists(role.Roles, rm) {
				role.Roles = append(role.Roles, rm)
			}
		}
		for _, rp := range validRolePrivileges {
			if !entryExists(role.Privileges, rp) {
				role.Privileges = append(role.Privileges, rp)
			}
		}
	} else if updateAction == "remove" {
		for _, rm := range validRoleMembers {
			if entryExists(role.Roles, rm) {
				role.Roles = removeEntryFromSlice(role.Roles, rm)
			}
		}
		for _, rp := range validRolePrivileges {
			if entryExists(role.Privileges, rp) {
				role.Privileges = removeEntryFromSlice(role.Privileges, rp)
			}
		}
	} else {
		log.Printf("%s : %s", getfuncName(), fmt.Sprintf(UpdateActionInvalidInfo, updateAction, UpdateActions))
		os.Exit(1)
	}

	if len(role.Roles)+len(role.Privileges) < 1 {
		log.Printf(roleItemsRequiredInfo, getfuncName())
		os.Exit(1)
	}

	if roleExists(id) {
		if roleExists(id) {
			payload, err := json.Marshal(m.Role{RoleID: id})
			logJsonMarshalError(err, jsonMarshalError)
			result := RunScript(deleteRoleScript, string(payload))
			if result.Status != successStatus {
				log.Printf("%s : %s", getfuncName(), setVerboseInfo)
			}
		} else {
			log.Printf(roleNotFoundInfo, id)
		}
		if !roleExists(id) {
			role := m.Role{RoleID: id, Name: id, Description: description, Source: getRoleSource(), Roles: role.Roles, Privileges: role.Privileges}
			payload, err := json.Marshal(role)
			logJsonMarshalError(err, jsonMarshalError)
			result := RunScript(createRoleScript, string(payload))
			if result.Status != successStatus {
				log.Printf("%s : %s", getfuncName(), setVerboseInfo)
			}
		}
		fmt.Printf(updateRoleSuccessInfo, id)
	} else {
		log.Printf(roleNotFoundInfo, id)
	}
}

func DeleteRole(id string) {
	if id == "" {
		log.Printf("%s : %s", getfuncName(), roleIDRequiredInfo)
		os.Exit(1)
	}
	if roleExists(id) {
		payload, err := json.Marshal(m.Role{RoleID: id})
		logJsonMarshalError(err, jsonMarshalError)
		result := RunScript(deleteRoleScript, string(payload))
		if result.Status == successStatus {
			log.Printf(deleteRoleSuccessInfo, id)
		} else {
			log.Printf("%s : %s", getfuncName(), setVerboseInfo)
		}
	} else {
		log.Printf(roleNotFoundInfo, id)
	}
}

func getRoles() []m.Role {
	payload, err := json.Marshal(m.Privilege{})
	logJsonMarshalError(err, getfuncName())
	result := RunScript(getRoleScript, string(payload))
	return result.Roles
}

func getRole(id string) m.Role {
	var role m.Role
	roles := getRoles()
	for _, r := range roles {
		if r.RoleID == id {
			role = r
		}
	}
	if role.RoleID == "" {
		log.Printf(roleNotFoundInfo, id)
		os.Exit(1)
	}
	return role
}

func getRoleIDs() []string {
	var rIDs []string
	roles := getRoles()
	for _, r := range roles {
		rIDs = append(rIDs, r.RoleID)
	}
	return rIDs
}

func roleExists(id string) bool {
	roles := getRoles()
	for _, r := range roles {
		if r.RoleID == id {
			return true
		}
	}
	return false
}

func getRoleDesc(description string) string {
	if description == "" {
		return defaultRoleDescription
	}
	return description
}

func getRoleSource() string {
	return defaultRoleSource
}

func validateRoleMembers(id, roleMembers string) []string {
	var validList []string
	roleMembersList := strings.Split(strings.Replace(roleMembers, " ", "", -1), ",")
	if len(roleMembersList) >= 1 && roleMembers != "" {
		for _, rm := range roleMembersList {
			if roleExists(rm) {
				if id != rm {
					validList = append(validList, rm)
				} else {
					log.Printf(cannotBeSameRoleInfo, rm, id)
				}
			} else {
				log.Printf(roleMemberNotFoundInfo, rm)
			}
		}
		if len(validList) < 1 {
			log.Println(noValidRoleMemberInfo)
		}
	} else {
		log.Println(noRoleMemberProvidedInfo)
	}
	return validList
}

func validateRolePrivileges(rolePrivileges string) []string {
	var validList []string
	rolePrivilegesList := strings.Split(strings.Replace(rolePrivileges, " ", "", -1), ",")
	if len(rolePrivilegesList) >= 1 && rolePrivileges != "" {
		for _, rp := range rolePrivilegesList {
			if privilegeExists(rp) {
				validList = append(validList, getPrivilegeID(rp))
			} else {
				log.Printf(rolePrivilegeNotFoundInfo, rp)
			}
		}
		if len(validList) < 1 {
			log.Println(noValidRolePrivilegeInfo)
		}
	} else {
		log.Println(noRolePrivilegesIProvidedInfo)
	}
	return validList
}
