package backend

import (
	m "com/privatesquare/go/nexus3-repository-cli/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func ListScripts(name string) {
	if name != "" {
		script := getScript(name)
		fmt.Println(script)
	} else {
		scriptsList := getScripts()
		sort.Strings(scriptsList)
		printStringSlice(scriptsList)
		if len(scriptsList) == 0 {
			fmt.Println("There are no scripts available in nexus")
		} else {
			fmt.Printf("No of scripts in nexus : %d\n", len(scriptsList))
		}
	}
}

func AddOrUpdateScript(name string) {
	if name == "" {
		log.Printf("%s : %s", getfuncName(), scriptNameRequiredInfo)
		os.Exit(1)
	}
	if !scriptExists(name) {
		AddScript(name)
	} else {
		UpdateScript(name)
	}
}

func AddScript(name string) {
	if name == "" {
		log.Printf("%s : %s", getfuncName(), scriptNameRequiredInfo)
		os.Exit(1)
	}
	url := fmt.Sprintf("%s/%s/%s", NexusURL, apiBase, scriptAPI)
	if !scriptExists(name) {
		payload, err := json.Marshal(m.Script{Name: name, Type: "groovy", Content: readFile(getScriptPath(name))})
		logJsonMarshalError(err, getfuncName())
		req := createBaseRequest("POST", url, m.RequestBody{Json: payload})
		_, status := httpRequest(req)
		if status == "204 No Content" {
			if Debug {
				log.Printf("The script %q is added to nexus\n", name)
			}
		} else {
			log.Printf("%s : %s", getfuncName(), setVerboseInfo)
			os.Exit(1)
		}
	} else {
		log.Printf("The script %q already exists in nexus\n", name)
	}
}

func UpdateScript(scriptName string) {
	if scriptName == "" {
		log.Printf("%s : %s", getfuncName(), scriptNameRequiredInfo)
		os.Exit(1)
	}
	url := fmt.Sprintf("%s/%s/%s/%s", NexusURL, apiBase, scriptAPI, scriptName)
	if scriptExists(scriptName) {
		payload, err := json.Marshal(m.Script{Name: scriptName, Type: "groovy", Content: readFile(getScriptPath(scriptName))})
		logJsonMarshalError(err, getfuncName())
		req := createBaseRequest("PUT", url, m.RequestBody{Json: payload})
		_, status := httpRequest(req)
		if status == "204 No Content" {
			if Debug {
				log.Printf("The script %q is updated in nexus\n", scriptName)
			}
		} else {
			log.Printf("%s : %s", getfuncName(), setVerboseInfo)
			os.Exit(1)
		}
	} else {
		log.Printf("The script %q does not exists in nexus\n", scriptName)
	}
}

func DeleteScript(scriptName string) {
	if scriptName == "" {
		log.Printf("%s : %s", getfuncName(), scriptNameRequiredInfo)
		os.Exit(1)
	}
	url := fmt.Sprintf("%s/%s/%s/%s", NexusURL, apiBase, scriptAPI, scriptName)
	if scriptExists(scriptName) {
		payload, err := json.Marshal(m.Script{Name: scriptName, Type: "groovy", Content: readFile(getScriptPath(scriptName))})
		logJsonMarshalError(err, getfuncName())
		req := createBaseRequest("Delete", url, m.RequestBody{Json: payload})
		_, status := httpRequest(req)
		if status == "204 No Content" {
			if Debug {
				log.Printf("The script %q is deleted from nexus\n", scriptName)
			}
		} else {
			log.Printf("%s : %s", getfuncName(), setVerboseInfo)
			os.Exit(1)
		}
	} else {
		log.Printf("The script %q does not exists in nexus\n", scriptName)
	}
}

func RunScript(scriptName, payload string) m.ScriptResult {
	if scriptName == "" {
		log.Printf("%s : %s", getfuncName(), scriptNameRequiredInfo)
		os.Exit(1)
	}
	var (
		output m.ScriptOutput
		result m.ScriptResult
	)
	AddOrUpdateScript(scriptName)
	url := fmt.Sprintf("%s/%s/%s/%s/run", NexusURL, apiBase, scriptAPI, scriptName)
	req := createBaseRequest("POST", url, m.RequestBody{Text: payload})
	respBody, status := httpRequest(req)
	if status == "200 OK" {
		if Debug {
			log.Printf("The script %q was executed successfully\n", scriptName)
		}
	} else {
		log.Printf("%s : %s", getfuncName(), setVerboseInfo)
		os.Exit(1)
	}
	err := json.Unmarshal(respBody, &output)
	logJsonUnmarshalError(err, getfuncName())
	err = json.Unmarshal([]byte(output.Result), &result)
	logJsonUnmarshalError(err, getfuncName())
	return result
}

func getScripts() []string {
	var (
		url         = fmt.Sprintf("%s/%s/%s", NexusURL, apiBase, scriptAPI)
		scripts     []m.Script
		scriptsList []string
	)
	req := createBaseRequest("GET", url, m.RequestBody{})
	respBody, status := httpRequest(req)
	if status != "200 OK" {
		log.Printf("%s : %s", getfuncName(), setVerboseInfo)
		os.Exit(1)
	} else {
		err := json.Unmarshal(respBody, &scripts)
		logJsonUnmarshalError(err, getfuncName())
		for _, s := range scripts {
			scriptsList = append(scriptsList, s.Name)
		}
	}
	return scriptsList
}

func getScript(scriptName string) m.Script {
	if scriptName == "" {
		log.Printf("%s : %s", getfuncName(), scriptNameRequiredInfo)
		os.Exit(1)
	}
	var (
		url    = fmt.Sprintf("%s/%s/%s/%s", NexusURL, apiBase, scriptAPI, scriptName)
		script m.Script
	)
	req := createBaseRequest("GET", url, m.RequestBody{})
	respBody, status := httpRequest(req)
	if status == "200 OK" {
		err := json.Unmarshal(respBody, &script)
		logJsonUnmarshalError(err, getfuncName())
	} else if status == "404 Not Found" {
		log.Printf("The script %q does not exist\n", scriptName)
		os.Exit(1)
	} else {
		log.Printf("%s : %s", getfuncName(), setVerboseInfo)
		os.Exit(1)
	}
	return script
}

func getScriptPath(scriptName string) string {
	if scriptName == "" {
		log.Printf("%s : %s", getfuncName(), scriptNameRequiredInfo)
		os.Exit(1)
	}
	return fmt.Sprintf("%s/%s.groovy", scriptBasePath, scriptName)
}

func scriptExists(scriptName string) bool {
	if scriptName == "" {
		log.Printf("%s : %s", getfuncName(), scriptNameRequiredInfo)
		os.Exit(1)
	}
	url := fmt.Sprintf("%s/%s/%s/%s", NexusURL, apiBase, scriptAPI, scriptName)
	req := createBaseRequest("GET", url, m.RequestBody{})
	_, status := httpRequest(req)
	if status == "200 OK" {
		return true
	}
	return false
}
