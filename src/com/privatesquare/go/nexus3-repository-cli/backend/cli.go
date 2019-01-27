package backend

import (
	m "com/privatesquare/go/nexus3-repository-cli/model"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func SetCLIConfiguration() {
	configuration := m.CLIConfiguration{NexusURL: NexusURL, Username: AuthUser.Username, Password: AuthUser.Password}
	configureJson, err := json.Marshal(configuration)
	logJsonMarshalError(err, jsonMarshalError)
	writeFile("nexus3-repository-cli.json", configureJson)
}

func getCLIConfiguration() m.CLIConfiguration {
	var conf m.CLIConfiguration
	data := readFile("nexus3-repository-cli.json")
	err := json.Unmarshal([]byte(data), &conf)
	logJsonUnmarshalError(err, jsonUnmarshalError)
	return conf
}

func SetConnectionDetails() {
	if fileExists("nexus3-repository-cli.json") {
		conf := getCLIConfiguration()
		NexusURL = conf.NexusURL
		AuthUser.Username = conf.Username
		AuthUser.Password = conf.Password
	} else {
		fmt.Println("Server connection details are not set...")
		fmt.Printf("Run %q to set the connection details.", "nexus3-repository-cli configure")
		os.Exit(1)
	}
}

func Usage() {
	flag.Usage = func() {
		fmt.Printf("Usage: nexus3-repository-cli [command]\n\n")
		fmt.Printf("[commands]\n  configurate\tSet nexus connection details\n  script  \tNexus script operations\n  repo  \tNexus repository operations\n\n")
	}
}

//TODO : Set usage as constants?
func ConfigureCommandUsage(fs *flag.FlagSet) {
	fs.Usage = func() {
		fmt.Printf("Usage: nexus3-repository-cli configure [args]\n\n")
		fmt.Println("[args]")
		fs.PrintDefaults()
		fmt.Println()
	}
}

//TODO : Set usage as constants?
func ScriptCommandUsage(fs *flag.FlagSet) {
	fs.Usage = func() {
		fmt.Printf("Usage: nexus3-repository-cli script [args] [options]\n\n")
		fmt.Printf("[args]\n  -task string\n\tScript Task (list|add|update|add-or-update|delete|run). (Required)\n  " +
			"-script-name string\n\tName of the script to be executed in nexus. The script should exist under the path ./scripts/groovy.\n  " +
			"-payload string\n\tArguments to be passed to a nexus script can be sent as a payload during script execution.\n\n" +
			"[options]\n  " +
			"-skip-tls\n\tSkip TLS verification for the nexus server instance.\n  " +
			"-debug\n\tSet Default for more information on the nexus script execution.\n  " +
			"-verbose\n\tSet Verbose for detailed http request and response logs.\n\n")
	}
}

//TODO : Set usage as constants?
func RepoCommandUsage(fs *flag.FlagSet) {
	fs.Usage = func() {
		fmt.Printf("Usage: nexus3-repository-cli repo [args] [options]\n\n")
		fmt.Printf("[args]\n  -repo-name string\n\tNexus repository name.\n  " +
			"-repo-format string\n\tRepository format (maven|npm|nuget|docker).\n  " +
			"-remote-url string\n\tRemote URL to be proxied in nexus.\n  " +
			"-repo-members string\n\tComma-separated repository names that should be added to a group repo.\n  " +
			"-release\n\tSet this flag to create a releases maven repository.\n" +
			"[options]\n  " +
			"-skip-tls\n\tSkip TLS verification for the nexus server instance.\n  " +
			"-debug\n\tSet Default for more information on the nexus script execution.\n  " +
			"-verbose\n\tSet Verbose for detailed http request and response logs.\n\n")
	}
}
