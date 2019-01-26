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

func CreateMavenHostedRepository(repoName, blobStoreName string, release bool) {
	if repoName == "" {
		log.Printf("%s : %s", getfuncName(), repoNameRequiredInfo)
		os.Exit(1)
	}
	payload, err := json.Marshal(m.Repository{Name: repoName, BlobStoreName: getBlobStoreName(blobStoreName), VersionPolicy: getVersionPolicy(release)})
	logJsonMarshalError(err, getfuncName())
	result := RunScript("create-maven-hosted", string(payload))
	printCreateRepoStatus(repoName, result.Status)
}

func CreateMavenProxyRepository(repoName, blobStoreName, remoteUrl string) {
	if repoName == "" {
		log.Printf("%s : %s", getfuncName(), repoNameRequiredInfo)
		os.Exit(1)
	}
	payload, err := json.Marshal(m.Repository{Name: repoName, RemoteURL: remoteUrl, BlobStoreName: getBlobStoreName(blobStoreName)})
	logJsonMarshalError(err, getfuncName())
	result := RunScript("create-maven-proxy", string(payload))
	printCreateRepoStatus(repoName, result.Status)
}

func CreateMavenGroupRepository(repoName, blobStoreName string, repoMembers []string) {
	if repoName == "" {
		log.Printf("%s : %s", getfuncName(), repoNameRequiredInfo)
		os.Exit(1)
	}
	payload, err := json.Marshal(m.Repository{Name: repoName, Members: repoMembers, BlobStoreName: getBlobStoreName(blobStoreName)})
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
		log.Printf("Repository %s was created\n", repoName)
	} else if status == "302 Found" {
		log.Printf("Repository %s already exists\n", repoName)
	} else {
		log.Printf("Error creating repository : %s\n", setVerboseInfo)
	}
}

func printDeleteRepoStatus(repoName, status string) {
	if status == "200 OK" {
		fmt.Printf("Repository %q was deleted\n", repoName)
	} else if status == "404 Not Found" {
		fmt.Printf("Repository %q was not found\n", repoName)
	} else {
		log.Printf("Error deleting repository : %s\n", setVerboseInfo)
	}
}
