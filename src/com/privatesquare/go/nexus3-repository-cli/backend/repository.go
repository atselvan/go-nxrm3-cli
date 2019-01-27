package backend

import (
	m "com/privatesquare/go/nexus3-repository-cli/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
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

func CreateMavenHostedRepository(repoName, blobStoreName string, release bool) {
	if repoName == "" {
		log.Printf("%s : %s", getfuncName(), mavenHostedRepoRequiredInfo)
		os.Exit(1)
	}
	payload, err := json.Marshal(m.Repository{Name: repoName, BlobStoreName: getBlobStoreName(blobStoreName), VersionPolicy: getVersionPolicy(release)})
	logJsonMarshalError(err, getfuncName())
	result := RunScript("create-maven-hosted", string(payload))
	printCreateRepoStatus(repoName, result.Status)
}

func CreateMavenProxyRepository(repoName, blobStoreName, remoteURL string) {
	if repoName == "" || remoteURL == "" {
		log.Printf("%s : %s", getfuncName(), proxyRepoRequiredInfo)
		os.Exit(1)
	}
	payload, err := json.Marshal(m.Repository{Name: repoName, RemoteURL: remoteURL, BlobStoreName: getBlobStoreName(blobStoreName)})
	logJsonMarshalError(err, getfuncName())
	result := RunScript("create-maven-proxy", string(payload))
	printCreateRepoStatus(repoName, result.Status)
}

func CreateMavenGroupRepository(repoName, blobStoreName string, repoMembers string) {
	if repoName == "" || repoMembers == "" {
		log.Printf("%s : %s", getfuncName(), groupRequiredInfo)
		os.Exit(1)
	}
	for _, r := range strings.Split(repoMembers, ",") {
		if !repositoryExists(r) {
			log.Printf("Repository %q does not exist in nexus. Please check the repo-members value and try again\n", r)
			os.Exit(1)
		}
	}
	payload, err := json.Marshal(m.Repository{Name: repoName, Members: strings.Split(repoMembers, ","), BlobStoreName: getBlobStoreName(blobStoreName)})
	logJsonMarshalError(err, getfuncName())
	result := RunScript("create-maven-group", string(payload))
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
		versionPolicy = "release"
	} else {
		versionPolicy = "snapshot"
	}
	return versionPolicy
}

func printCreateRepoStatus(repoName, status string) {
	if status == "200 OK" {
		log.Printf("Repository %s was created in nexus\n", repoName)
	} else if status == "302 Found" {
		log.Printf("Repository %s already exists in nexus\n", repoName)
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
