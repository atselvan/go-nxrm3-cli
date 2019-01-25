package backend

import (
	m "com/privatesquare/go/nexus3-repository-cli/model"
	"encoding/json"
	"fmt"
	"log"
)

func ListScripts(){
	scriptsList := getScripts()
	for _, s := range scriptsList {
		fmt.Println(s)
	}
	fmt.Printf("No of scripts in nexus : %d\n", len(scriptsList))
}

func AddOrUpdateScript(scriptName string){
	if !scriptExists(scriptName){
		AddScript(scriptName)
	} else {
		UpdateScript(scriptName)
	}
}

func AddScript(scriptName string){
	url := fmt.Sprintf("%s/%s/%s", NexusURL, apiBase, scriptAPI)
	if !scriptExists(scriptName) {
		payload, err := json.Marshal(m.Script{Name: scriptName, Type: "groovy", Content: readFile(getScriptPath(scriptName))})
		logError(err, "Add Script : Json Marshal Error")
		req := createBaseRequest("POST", url, payload)
		_, status := httpRequest(req)
		if status == "204 No Content" {
			log.Printf("Added the script %q in nexus\n", scriptName)
		}
	} else {
		log.Printf("Script %q already exists in nexus\n", scriptName)
	}
}

func UpdateScript(scriptName string){
	url := fmt.Sprintf("%s/%s/%s/%s", NexusURL, apiBase, scriptAPI, scriptName)
	if scriptExists(scriptName) {
		payload, err := json.Marshal(m.Script{Name: scriptName, Type: "groovy", Content: readFile(getScriptPath(scriptName))})
		logError(err, "Update Script : Json Marshal Error")
		req := createBaseRequest("PUT", url, payload)
		_, status := httpRequest(req)
		if status == "204 No Content" {
			log.Printf("Updated the script %q in nexus\n", scriptName)
		}
	} else {
		log.Printf("Script %q does not exists in nexus\n", scriptName)
	}
}

func DeleteScript(scriptName string){
	url := fmt.Sprintf("%s/%s/%s/%s", NexusURL, apiBase, scriptAPI, scriptName)
	if scriptExists(scriptName) {
		payload, err := json.Marshal(m.Script{Name: scriptName, Type: "groovy", Content: readFile(getScriptPath(scriptName))})
		logError(err, "Delete Script : Json Marshal Error")
		req := createBaseRequest("Delete", url, payload)
		_, status := httpRequest(req)
		if status == "204 No Content" {
			log.Printf("Deleted the script %q from nexus\n", scriptName)
		}
	} else {
		log.Printf("Script %q does not exists in nexus\n", scriptName)
	}
}

func RunScript(scriptName string){
	url := fmt.Sprintf("%s/%s/%s/%s/run", NexusURL, apiBase, scriptAPI, scriptName)

	payload := `{
    			"name": "test-repo-10"
				}`
	
	postreq := createBaseRequest1("POST", url, payload)

	respBody, status := httpRequest(postreq)

	fmt.Println(string(respBody), status)
}

func getScripts() []string{
	var (
		url = fmt.Sprintf("%s/%s/%s", NexusURL, apiBase, scriptAPI)
		scripts []m.Script
		scriptsList []string
	)
	req := createBaseRequest("GET", url, nil)
	respBody, status := httpRequest(req)
	err := json.Unmarshal(respBody, &scripts)
	logError(err, "Get Scripts : JSON Unmarshal error")
	for _, s := range scripts{
		scriptsList = append(scriptsList, s.Name)
	}
	if status != "200 OK"{
		log.Fatal("Get Scripts : Error getting scripts list. Set verbose for more information")
	}
	return scriptsList
}

func GetScript(scriptName string) m.Script{
	var (
		url = fmt.Sprintf("%s/%s/%s/%s", NexusURL, apiBase, scriptAPI, scriptName)
		script m.Script
	)
	req := createBaseRequest("GET", url, nil)
	respBody, status := httpRequest(req)
	err := json.Unmarshal(respBody, &script)
	logError(err, "Get Script : JSON Unmarshal error")
	if status != "200 OK"{
		log.Fatal("Get Script : Error getting the script. Set verbose for more information")
	}
	return script
} 

func getScriptPath(scriptName string) string{
	return fmt.Sprintf("%s/%s.groovy", scriptBasePath, scriptName)
}

func scriptExists(scriptName string) bool{
	url := fmt.Sprintf("%s/%s/%s/%s", NexusURL, apiBase, scriptAPI, scriptName)
	req := createBaseRequest("GET", url, nil)
	_, status := httpRequest(req)
	if status == "200 OK"{
		return true
	}
	return false
}
