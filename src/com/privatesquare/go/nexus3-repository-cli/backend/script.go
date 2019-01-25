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
	if len(scriptsList) == 0 {
		fmt.Println("There are no scripts available in nexus")
	} else {
		fmt.Printf("No of scripts in nexus : %d\n", len(scriptsList))
	}
}

func AddOrUpdateScript(scriptName string){
	if scriptName == "" {
		log.Fatal("Add/Update Script Error : scriptName is a required parameter")
	}
	if !scriptExists(scriptName){
		AddScript(scriptName)
	} else {
		UpdateScript(scriptName)
	}
}

func AddScript(scriptName string){
	if scriptName == "" {
		log.Fatal("Add Script Error : scriptName is a required parameter")
	}
	url := fmt.Sprintf("%s/%s/%s", NexusURL, apiBase, scriptAPI)
	if !scriptExists(scriptName) {
		payload, err := json.Marshal(m.Script{Name: scriptName, Type: "groovy", Content: readFile(getScriptPath(scriptName))})
		logError(err, "Add Script : Json Marshal Error")
		req := createBaseRequest("POST", url, payload)
		_, status := httpRequest(req)
		if status == "204 No Content" {
			log.Printf("Added the script %q in nexus\n", scriptName)
		} else {
			log.Printf("Add Script Error : Set verbose flag for more information")
		}
	} else {
		log.Printf("Script %q already exists in nexus\n", scriptName)
	}
}

func UpdateScript(scriptName string){
	if scriptName == "" {
		log.Fatal("Update Script Error : scriptName is a required parameter")
	}
	url := fmt.Sprintf("%s/%s/%s/%s", NexusURL, apiBase, scriptAPI, scriptName)
	if scriptExists(scriptName) {
		payload, err := json.Marshal(m.Script{Name: scriptName, Type: "groovy", Content: readFile(getScriptPath(scriptName))})
		logError(err, "Update Script : Json Marshal Error")
		req := createBaseRequest("PUT", url, payload)
		_, status := httpRequest(req)
		if status == "204 No Content" {
			log.Printf("Updated the script %q in nexus\n", scriptName)
		} else {
			log.Printf("Update Script Error : Set verbose flag for more information")
		}
	} else {
		log.Printf("Script %q does not exists in nexus\n", scriptName)
	}
}

func DeleteScript(scriptName string){
	if scriptName == "" {
		log.Fatal("Delete Script Error : scriptName is a required parameter")
	}
	url := fmt.Sprintf("%s/%s/%s/%s", NexusURL, apiBase, scriptAPI, scriptName)
	if scriptExists(scriptName) {
		payload, err := json.Marshal(m.Script{Name: scriptName, Type: "groovy", Content: readFile(getScriptPath(scriptName))})
		logError(err, "Delete Script : Json Marshal Error")
		req := createBaseRequest("Delete", url, payload)
		_, status := httpRequest(req)
		if status == "204 No Content" {
			log.Printf("Deleted the script %q from nexus\n", scriptName)
		} else {
			log.Printf("Delete Script Error : Set verbose flag for more information")
		}
	} else {
		log.Printf("Script %q does not exists in nexus\n", scriptName)
	}
}

func RunScript(scriptName, payload string) []byte{
	if scriptName == "" {
		log.Fatal("Run Script Error : scriptName is a required parameter")
	}
	AddOrUpdateScript(scriptName)
	url := fmt.Sprintf("%s/%s/%s/%s/run", NexusURL, apiBase, scriptAPI, scriptName)
	req := createBaseRequest1("POST", url, payload)
	respBody, status := httpRequest(req)
	if status == "200 OK"{
		log.Printf("Script %q was executed successfully\n", scriptName)
	} else {
		log.Printf("Run Script Error : Set verbose flag for more information")
	}
	return respBody
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
		log.Fatal("Get Scripts Error : Set verbose flag for more information")
	}
	return scriptsList
}

func GetScript(scriptName string) m.Script{
	if scriptName == "" {
		log.Fatal("Get Script Error : scriptName is a required parameter")
	}
	var (
		url = fmt.Sprintf("%s/%s/%s/%s", NexusURL, apiBase, scriptAPI, scriptName)
		script m.Script
	)
	req := createBaseRequest("GET", url, nil)
	respBody, status := httpRequest(req)
	err := json.Unmarshal(respBody, &script)
	logError(err, "Get Script : JSON Unmarshal error")
	if status != "200 OK"{
		log.Fatal("Get Script Error : Set verbose flag for more information")
	}
	return script
} 

func getScriptPath(scriptName string) string{
	if scriptName == "" {
		log.Fatal("Get Script Path Error : scriptName is a required parameter")
	}
	return fmt.Sprintf("%s/%s.groovy", scriptBasePath, scriptName)
}

func scriptExists(scriptName string) bool{
	if scriptName == "" {
		log.Fatal("Script Exists Error : scriptName is a required parameter")
	}
	url := fmt.Sprintf("%s/%s/%s/%s", NexusURL, apiBase, scriptAPI, scriptName)
	req := createBaseRequest("GET", url, nil)
	_, status := httpRequest(req)
	if status == "200 OK"{
		return true
	}
	return false
}
