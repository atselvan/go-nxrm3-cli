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

func ListRepositories(name, format string) {
	var repositoryList []string

	if name != "" {
		repository := getRepository(name)
		fmt.Printf("Name: %s\nRecipe: %s\nURL: %s\n", repository.Name, repository.Recipe, repository.URL)
	} else if name == "" && format == "" {
		repositoryList = getRepositoryList()
	} else {
		repositoryList = getRepositoryListByFormat(validateRepositoryFormat(name))
	}
	if name == "" {
		printStringSlice(repositoryList)
		fmt.Printf("Number of repositories : %d\n", len(repositoryList))
	}
}

func CreateHosted(name, blobStoreName, format string, dockerHttpPort, dockerHttpsPort int, releases bool) {
	if name == "" || format == "" {
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

	repository := m.Repository{Name: name, Format: format, Recipe: recipe, Attributes: attributes}
	payload, err := json.Marshal(repository)
	logJsonMarshalError(err, getfuncName())
	result := RunScript(createHostedRepoScript, string(payload))
	printCreateRepoStatus(name, result.Status)
}

func CreateProxy(name, blobStoreName, format, remoteURL, proxyUsername, proxyPassword string, dockerHttpPort, dockerHttpsPort int, releases bool) {
	if name == "" || remoteURL == "" || format == "" {
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

	repository := m.Repository{Name: name, Format: format, Recipe: recipe, Attributes: attributes}
	payload, err := json.Marshal(repository)
	logJsonMarshalError(err, getfuncName())
	result := RunScript(createProxyRepoScript, string(payload))
	printCreateRepoStatus(name, result.Status)
}

func CreateGroup(name, blobStoreName, format, repoMembers string, dockerHttpPort, dockerHttpsPort int, releases bool) {
	if name == "" || repoMembers == "" || format == "" {
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

	repository := m.Repository{Name: name, Format: format, Recipe: recipe, Attributes: attributes}
	payload, err := json.Marshal(repository)
	logJsonMarshalError(err, getfuncName())
	result := RunScript(createGroupRepoScript, string(payload))
	printCreateRepoStatus(name, result.Status)
}

func AddMembersToGroup(name, format, repoMembers string) {
	if name == "" || repoMembers == "" || format == "" {
		log.Printf("%s : %s", getfuncName(), groupRequiredInfo)
		os.Exit(1)
	}
	if repositoryExists(name) {
		repo := getRepository(name)
		validateGroupRepo(repo)
		format = validateRepositoryFormat(format)
		validList := validateGroupMembers(repoMembers, format)
		currentMembers := repo.Attributes.Group.MemberNames
		for _, newMember := range validList {
			if entryExists(currentMembers, newMember) {
				log.Printf(groupMemberAlreadyExistsInfo, newMember, name)
			} else if newMember == name {
				log.Printf(cannotBeSameRepoInfo, newMember, name)
			} else {
				log.Printf(groupMemberAddSuccessInfo, newMember, name)
				currentMembers = append(currentMembers, newMember)
			}
		}
		repo.Attributes.Group = m.Group{MemberNames: currentMembers}
		repository := m.Repository{Name: name, Format: format, Attributes: repo.Attributes}
		payload, err := json.Marshal(repository)
		logJsonMarshalError(err, getfuncName())
		result := RunScript(updateGroupMembersScript, string(payload))
		printUpdateRepoStatus(name, result.Status)
	} else {
		log.Printf(repositoryNotFoundInfo, name)
	}
}

func RemoveMembersFromGroup(name, format, repoMembers string) {
	if name == "" || repoMembers == "" || format == "" {
		log.Printf("%s : %s", getfuncName(), groupRequiredInfo)
		os.Exit(1)
	}
	if repositoryExists(name) {
		repo := getRepository(name)
		validateGroupRepo(repo)
		format = validateRepositoryFormat(format)
		validList := validateGroupMembers(repoMembers, format)
		currentMembers := repo.Attributes.Group.MemberNames
		for _, newMember := range validList {
			if !entryExists(currentMembers, newMember) {
				log.Printf(groupMemberRemoveNotFoundInfo, newMember, name)
			} else if newMember == name {
				log.Printf(cannotBeSameRepoInfo, newMember, name)
			} else {
				log.Printf(groupMemberRemoveSuccessInfo, newMember, name)
				currentMembers = removeEntryFromSlice(currentMembers, newMember)
			}
		}
		repo.Attributes.Group = m.Group{MemberNames: currentMembers}
		repository := m.Repository{Name: name, Format: format, Attributes: repo.Attributes}
		payload, err := json.Marshal(repository)
		logJsonMarshalError(err, getfuncName())
		result := RunScript(updateGroupMembersScript, string(payload))
		printUpdateRepoStatus(name, result.Status)
	} else {
		log.Printf(repositoryNotFoundInfo, name)
	}
}

func DeleteRepository(name string) {
	if name == "" {
		log.Printf("%s : %s", getfuncName(), repoNameRequiredInfo)
		os.Exit(1)
	}
	payload, err := json.Marshal(m.Repository{Name: name})
	logJsonMarshalError(err, getfuncName())
	result := RunScript(deleteRepoScript, string(payload))
	printDeleteRepoStatus(name, result.Status)
}

func getRepository(name string) m.Repository {
	if name == "" {
		log.Printf("%s : %s", getfuncName(), repoNameRequiredInfo)
		os.Exit(1)
	}
	payload, err := json.Marshal(m.Repository{Name: name})
	logJsonMarshalError(err, getfuncName())
	result := RunScript(getRepoScript, string(payload))
	if result.Status != successStatus {
		log.Printf("%s : %s", getfuncName(), setVerboseInfo)
		os.Exit(1)
	}
	return m.Repository{Name: result.Name, URL: result.URL, Type: result.Type, Format: result.Format, Recipe: result.Recipe, Attributes: result.Attributes}
}

func getRepositories() []m.Repository {
	url := fmt.Sprintf("%s/%s/%s", NexusURL, apiBase, repositoryPath)
	var repositories []m.Repository
	req := createBaseRequest("GET", url, m.RequestBody{})
	respBody, status := httpRequest(req)
	if status != successStatus {
		log.Printf("%s : %s", getfuncName(), setVerboseInfo)
		os.Exit(1)
	} else {
		err := json.Unmarshal(respBody, &repositories)
		logJsonUnmarshalError(err, getfuncName())
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

func getRepositoryListByFormat(format string) []string {
	var repositoryList []string
	repositories := getRepositories()
	for _, r := range repositories {
		if format == r.Format {
			repositoryList = append(repositoryList, r.Name)
		}
	}
	return repositoryList
}

func repositoryExists(name string) bool {
	if name == "" {
		log.Printf("%s : %s", getfuncName(), repoNameRequiredInfo)
		os.Exit(1)
	}
	var isExists bool
	payload, err := json.Marshal(m.Repository{Name: name})
	logJsonMarshalError(err, getfuncName())
	result := RunScript(getRepoScript, string(payload))
	if result.Status == successStatus {
		isExists = true
	} else if result.Status == notFoundStatus {
		isExists = false
	} else {
		log.Printf("%s : %s", getfuncName(), setVerboseInfo)
		os.Exit(1)
	}
	return isExists
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
		log.Printf("%s : %s", getfuncName(), fmt.Sprintf(RepoFormatNotValidInfo, format, RepoFormats))
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
		log.Printf("%s : %s\n", getfuncName(), proxyCredsNotValidInfo)
		os.Exit(1)
	}
}

func validateRemoteURL(url string) {
	httpRegex, _ := regexp2.Compile(`^(http://).*`)
	httpsRegex, _ := regexp2.Compile(`^(https://).*`)
	if httpRegex.MatchString(url) || httpsRegex.MatchString(url) {
		return
	} else {
		log.Printf("%s : %s", getfuncName(), fmt.Sprintf(remoteURLNotValidInfo, url))
		os.Exit(1)
	}
}

func validateGroupRepo(repo m.Repository){
	if strings.Contains(repo.Recipe, "group"){
		return
	}else {
		log.Printf(notAGroupRepoInfo, repo.Name)
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
				log.Printf(groupMemberInvalidFormatInfo, repoMember, format)
			}
		} else {
			log.Printf(groupMemberNotFoundInfo, repoMember)
		}
	}
	if len(validList) < 1 {
		log.Printf("%s : %s\n", getfuncName(), groupMemberRequiredInfo)
		os.Exit(1)
	}
	return validList
}

func printCreateRepoStatus(name, status string) {
	if status == successStatus {
		log.Printf( repoCreatedInfo, name)
	} else if status == foundStatus {
		log.Printf(repoExistsInfo, name)
	} else {
		log.Printf(repoCreateErrorInfo, setVerboseInfo)
	}
}

func printUpdateRepoStatus(name, status string) {
	if status == successStatus {
		log.Printf(repoUpdatedStatus, name)
	} else if status == notFoundStatus {
		log.Printf(repositoryNotFoundInfo, name)
	} else {
		log.Printf(repoUpdateErrorInfo, setVerboseInfo)
	}
}

func printDeleteRepoStatus(name, status string) {
	if status == successStatus {
		log.Printf(repoDeletedInfo, name)
	} else if status == notFoundStatus {
		log.Printf(repositoryNotFoundInfo, name)
	} else {
		log.Printf(repoDeleteErrorInfo, setVerboseInfo)
	}
}
