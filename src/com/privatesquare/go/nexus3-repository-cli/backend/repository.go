package backend

import (
	m "com/privatesquare/go/nexus3-repository-cli/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ListRepositories(repoName, repoFormat string) {
	var repositoryList []string

	if repoName != "" {
		repository := getRepository(repoName)
		fmt.Printf("Name: %s\nRecipe: %s\nURL: %s\n", repository.Name, repository.Recipe, repository.URL)
	} else if repoName == "" && repoFormat == "" {
		repositoryList = getRepositoryList()
	} else {
		repositoryList = getRepositoryListByFormat(repoFormat)
	}
	if repoName == "" {
		printStringSlice(repositoryList)
		fmt.Printf("Number of repositories : %d\n", len(repositoryList))
	}
}

func GetRepositoryAttributes(repoName string) {
	if repoName == "" {
		log.Printf("%s : %s", getfuncName(), repoNameRequiredInfo)
		os.Exit(1)
	}
	attribute := m.Attributes{Maven: m.Maven{VersionPolicy: "Releases", LayoutPolicy: "Something"}}
	payload, err := json.Marshal(attribute)
	logJsonMarshalError(err, getfuncName())
	SkipTLSVerification = true
	Verbose = true
	result := RunScript("get-repo-attributes", string(payload))
	fmt.Println(result)
}

func CreateHosted(repoName, blobStoreName, format string, dockerHttpPort, dockerHttpsPort int, releases bool) {
	if repoName == "" {
		log.Printf("%s : %s", getfuncName(), repoNameRequiredInfo)
		os.Exit(1)
	}
	format = validateRepositoryFormat(format)
	var attributes m.Attributes
	recipe := fmt.Sprintf("%s-hosted", format)
	if format == "maven2" {
		storage := m.Storage{BlobStoreName: getBlobStoreName(blobStoreName), StrictContentTypeValidation: true, WritePolicy: getWritePolicy(releases)}
		maven := m.Maven{VersionPolicy: getVersionPolicy(releases), LayoutPolicy: "STRICT"}
		attributes = m.Attributes{Storage: storage, Maven: maven}
	} else if format == "docker" {
		if dockerHttpPort == 0 && dockerHttpsPort == 0 {
			log.Printf("%s : %s", getfuncName(), dockerPortsInfo)
			os.Exit(1)
		}
		storage := m.Storage{BlobStoreName: getBlobStoreName(blobStoreName), StrictContentTypeValidation: true, WritePolicy: getWritePolicy(releases)}
		docker := m.Docker{HTTPPort: dockerHttpPort, HTTPSPort: dockerHttpsPort, ForceBasicAuth: true, V1Enabled: false}
		attributes = m.Attributes{Storage: storage, Docker: docker}
	} else {
		storage := m.Storage{BlobStoreName: getBlobStoreName(blobStoreName), StrictContentTypeValidation: true, WritePolicy: getWritePolicy(releases)}
		attributes = m.Attributes{Storage: storage}
	}
	repository := m.Repository{Name: repoName, Format: format, Recipe: recipe, Attribute: attributes}
	payload, err := json.Marshal(repository)
	logJsonMarshalError(err, getfuncName())
	result := RunScript("create-hosted-repo", string(payload))
	printCreateRepoStatus(repoName, result.Status)
}

func DeleteRepository(repoName string) {
	if repoName == "" {
		log.Printf("%s : %s", getfuncName(), repoNameRequiredInfo)
		os.Exit(1)
	}
	payload, err := json.Marshal(m.Repository{Name: repoName})
	logJsonMarshalError(err, getfuncName())
	result := RunScript("delete-repo", string(payload))
	printDeleteRepoStatus(repoName, result.Status)
}

func getRepository(repoName string) m.Repository {
	if repoName == "" {
		log.Printf("%s : %s", getfuncName(), repoNameRequiredInfo)
		os.Exit(1)
	}
	payload, err := json.Marshal(m.Repository{Name: repoName})
	logJsonMarshalError(err, getfuncName())
	result := RunScript("get-repo", string(payload))
	if result.Status != "200 OK" {
		log.Printf("%s : %s", getfuncName(), setVerboseInfo)
		os.Exit(1)
	}
	return m.Repository{Name: result.Name, Recipe: result.Recipe, URL: result.URL}
}

func getRepositories() []m.Repository {
	url := fmt.Sprintf("%s/%s/%s", NexusURL, apiBase, repositoryPath)
	var repositories []m.Repository
	req := createBaseRequest("GET", url, m.RequestBody{})
	respBody, status := httpRequest(req)
	err := json.Unmarshal(respBody, &repositories)
	logError(err, "Get Repositories : JSON Unmarshal Error")
	if status != "200 OK" {
		log.Printf("%s : %s", getfuncName(), setVerboseInfo)
		os.Exit(1)
	}
	return repositories
}

func getRepositoryList() []string {
	var repositoryList []string
	repositories := getRepositories()
	for _, r := range repositories {
		repositoryList = append(repositoryList, r.Name)
	}
	return repositoryList
}

func repositoryExists(repoName string) bool {
	if repoName == "" {
		log.Printf("%s : %s", getfuncName(), repoNameRequiredInfo)
		os.Exit(1)
	}
	var isExists bool
	payload, err := json.Marshal(m.Repository{Name: repoName})
	logJsonMarshalError(err, getfuncName())
	result := RunScript("get-repo", string(payload))
	if result.Status == "200 OK" {
		isExists = true
	} else if result.Status == "404 Not Found" {
		isExists = false
	} else {
		log.Printf("%s : %s", getfuncName(), setVerboseInfo)
		os.Exit(1)
	}
	return isExists
}

func getRepositoryListByFormat(repoFormat string) []string {
	var repositoryList []string
	repositories := getRepositories()
	for _, r := range repositories {
		if repoFormat == r.Format {
			repositoryList = append(repositoryList, r.Name)
		}
	}
	return repositoryList
}

func getBlobStoreName(blobStoreName string) string {
	if blobStoreName == "" {
		blobStoreName = "default"
	}
	return blobStoreName
}

func getVersionPolicy(release bool) string {
	var versionPolicy string
	if release {
		versionPolicy = "RELEASE"
	} else {
		versionPolicy = "SNAPSHOT"
	}
	return versionPolicy
}

func getWritePolicy(releases bool) string {
	var writePolicy string
	if releases {
		writePolicy = "ALLOW_ONCE"
	} else {
		writePolicy = "ALLOW"
	}
	return writePolicy
}

func validateRepositoryFormat(format string) string {
	if format == "" {
		log.Printf("%s : %s", getfuncName(), repoFormatRequiredInfo)
		os.Exit(1)
	}
	formatChoice := map[string]bool{"": true}
	for _, repoFormat := range RepoFormats {
		formatChoice[repoFormat] = true
	}
	if _, validChoice := formatChoice[format]; !validChoice {
		log.Printf("%q is not a valid repository format. Available repository formats are : %v\n", format, RepoFormats)
		os.Exit(1)
	}
	if format == "maven" {
		return "maven2"
	}
	return format
}

func printCreateRepoStatus(repoName, status string) {
	if status == "200 OK" {
		log.Printf("Repository %q was created in nexus\n", repoName)
	} else if status == "302 Found" {
		log.Printf("Repository %q already exists in nexus\n", repoName)
	} else {
		log.Printf("Error creating repository : %s\n", setVerboseInfo)
	}
}

func printDeleteRepoStatus(repoName, status string) {
	if status == "200 OK" {
		log.Printf("Repository %q was deleted from nexus\n", repoName)
	} else if status == "404 Not Found" {
		log.Printf("Repository %q was not found in nexus\n", repoName)
	} else {
		log.Printf("Error deleting repository : %s\n", setVerboseInfo)
	}
}
