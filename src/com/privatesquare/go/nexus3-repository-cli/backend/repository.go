package backend

import (
	m "com/privatesquare/go/nexus3-repository-cli/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	regexp2 "regexp"
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
	if repoName == "" || format == "" {
		log.Printf("%s : %s", getfuncName(), hostedRepoRequiredInfo)
		os.Exit(1)
	}
	format = validateRepositoryFormat(format)

	var attributes m.Attributes
	recipe := fmt.Sprintf("%s-hosted", format)
	storage := m.Storage{BlobStoreName: getBlobStoreName(blobStoreName), StrictContentTypeValidation: true, WritePolicy: getWritePolicy(releases)}

	if format == "maven2" {
		maven := m.Maven{VersionPolicy: getVersionPolicy(releases), LayoutPolicy: "STRICT"}
		attributes = m.Attributes{Storage: storage, Maven: maven}
	} else if format == "docker" {
		if dockerHttpPort == 0 && dockerHttpsPort == 0 {
			log.Printf("%s : %s", getfuncName(), dockerPortsInfo)
			os.Exit(1)
		}
		docker := m.Docker{HTTPPort: dockerHttpPort, HTTPSPort: dockerHttpsPort, ForceBasicAuth: true, V1Enabled: false}
		attributes = m.Attributes{Storage: storage, Docker: docker}
	} else {
		attributes = m.Attributes{Storage: storage}
	}

	repository := m.Repository{Name: repoName, Format: format, Recipe: recipe, Attribute: attributes}
	payload, err := json.Marshal(repository)
	logJsonMarshalError(err, getfuncName())
	result := RunScript("create-hosted-repo", string(payload))
	printCreateRepoStatus(repoName, result.Status)
}

func CreateProxy(repoName, blobStoreName, format, remoteURL, proxyUsername, proxyPassword string, dockerHttpPort, dockerHttpsPort int, releases bool) {
	if repoName == "" || remoteURL == "" || format == "" {
		log.Printf("%s : %s", getfuncName(), proxyRepoRequiredInfo)
		os.Exit(1)
	}
	format = validateRepositoryFormat(format)
	validateProxyAuthInfo(proxyUsername, proxyPassword)
	validateRemoteURL(remoteURL)

	var attributes m.Attributes
	recipe := fmt.Sprintf("%s-proxy", format)
	storage := m.Storage{BlobStoreName: getBlobStoreName(blobStoreName), StrictContentTypeValidation: true, WritePolicy: getWritePolicy(releases)}
	proxy := m.Proxy{RemoteURL: remoteURL, ContentMaxAge: -1, MetadataMaxAge: 1440}
	proxyAuth := m.HttpClientAuth{Username: proxyUsername, Password: proxyPassword}
	proxyHttpClient := m.HttpClient{Blocked: false, AutoBlock: true, Authentication: proxyAuth}
	negetiveCache := m.NegetiveCache{Enabled: true, TimeToLive: 1440}

	if format == "maven2" {
		maven := m.Maven{VersionPolicy: getVersionPolicy(releases), LayoutPolicy: "STRICT"}
		attributes = m.Attributes{Storage: storage, Maven: maven, Proxy: proxy, Httpclient: proxyHttpClient, NegativeCache: negetiveCache}
	} else if format == "docker" {
		if dockerHttpPort == 0 && dockerHttpsPort == 0 {
			log.Printf("%s : %s", getfuncName(), dockerPortsInfo)
			os.Exit(1)
		}
		docker := m.Docker{HTTPPort: dockerHttpPort, HTTPSPort: dockerHttpsPort, ForceBasicAuth: true, V1Enabled: false}
		dockerProxy := m.DockerProxy{IndexType: "REGISTRY"}
		attributes = m.Attributes{Storage: storage, Docker: docker, Proxy: proxy, DockerProxy: dockerProxy, Httpclient: proxyHttpClient, NegativeCache: negetiveCache}
	} else {
		storage := m.Storage{BlobStoreName: getBlobStoreName(blobStoreName), StrictContentTypeValidation: true, WritePolicy: getWritePolicy(releases)}
		attributes = m.Attributes{Storage: storage, Proxy: proxy, Httpclient: proxyHttpClient, NegativeCache: negetiveCache}
	}

	repository := m.Repository{Name: repoName, Format: format, Recipe: recipe, Attribute: attributes}
	payload, err := json.Marshal(repository)
	logJsonMarshalError(err, getfuncName())
	result := RunScript("create-proxy-repo", string(payload))
	printCreateRepoStatus(repoName, result.Status)
}

func CreateGroup(repoName, blobStoreName, format, repoMembers string, dockerHttpPort, dockerHttpsPort int, releases bool) {
	if repoName == "" || repoMembers == "" || format == "" {
		log.Printf("%s : %s", getfuncName(), groupRequiredInfo)
		os.Exit(1)
	}
	format = validateRepositoryFormat(format)
	validList := validateGroupMembers(repoMembers, format)

	var attributes m.Attributes
	recipe := fmt.Sprintf("%s-group", format)
	storage := m.Storage{BlobStoreName: getBlobStoreName(blobStoreName), StrictContentTypeValidation: true, WritePolicy: getWritePolicy(releases)}
	group := m.Group{MemberNames: validList}

	if format == "maven2" {
		maven := m.Maven{VersionPolicy: getVersionPolicy(releases), LayoutPolicy: "STRICT"}
		attributes = m.Attributes{Storage: storage, Maven: maven, Group: group}
	} else if format == "docker" {
		if dockerHttpPort == 0 && dockerHttpsPort == 0 {
			log.Printf("%s : %s", getfuncName(), dockerPortsInfo)
			os.Exit(1)
		}
		docker := m.Docker{HTTPPort: dockerHttpPort, HTTPSPort: dockerHttpsPort, ForceBasicAuth: true, V1Enabled: false}
		attributes = m.Attributes{Storage: storage, Docker: docker, Group: group}
	} else {
		storage := m.Storage{BlobStoreName: getBlobStoreName(blobStoreName), StrictContentTypeValidation: true, WritePolicy: getWritePolicy(releases)}
		attributes = m.Attributes{Storage: storage, Group: group}
	}

	repository := m.Repository{Name: repoName, Format: format, Recipe: recipe, Attribute: attributes}
	payload, err := json.Marshal(repository)
	logJsonMarshalError(err, getfuncName())
	result := RunScript("create-group-repo", string(payload))
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
		log.Printf("%s : %q is not a valid repository format. Available repository formats are : %v\n", getfuncName(), format, RepoFormats)
		os.Exit(1)
	}
	if format == "maven" {
		return "maven2"
	}
	return format
}

func validateProxyAuthInfo(proxyUsername, proxyPassword string) {
	if proxyUsername == "" && proxyPassword == "" {
		return
	} else if proxyUsername != "" && proxyPassword != "" {
		return
	} else {
		log.Printf("%s : You need to provide both proxy-user and proxy-pass to set credentials to a proxy repository\n", getfuncName())
		os.Exit(1)
	}
}

func validateRemoteURL(url string) {
	httpRegex, _ := regexp2.Compile(`^(http://).*`)
	httpsRegex, _ := regexp2.Compile(`^(https://).*`)
	if httpRegex.MatchString(url) || httpsRegex.MatchString(url) {
		return
	} else {
		log.Printf("%s : %q is an invalid url. URL must begin with either http:// or https://\n", getfuncName(), url)
		os.Exit(1)
	}
}

func validateGroupMembers(repoMembers, format string) []string {
	var validList []string
	repoMembersList := strings.Split(strings.Replace(repoMembers, " ", "", -1), ",")
	for _, repoMember := range repoMembersList {
		if repositoryExists(repoMember) {
			repoDetails := getRepository(repoMember)
			if strings.Contains(repoDetails.Recipe, format) {
				validList = append(validList, repoMember)
			} else {
				log.Printf("Repository %q is not a %q format repository, hence it cannot be added to the group repository\n", repoMember, format)
			}
		} else {
			log.Printf("Repository %q was not found, hence it cannot be added to the group repository\n", repoMember)
		}
	}
	if len(validList) < 1 {
		log.Printf("%s : Atleast one valid group member should be provided to create a group repository", getfuncName())
		os.Exit(1)
	}
	return validList
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
