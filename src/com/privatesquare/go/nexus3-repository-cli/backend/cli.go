package backend

import (
	m "com/privatesquare/go/nexus3-repository-cli/model"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func StoreConnectionDetails() {
	configuration := m.ConnDetails{NexusURL: NexusURL, Username: AuthUser.Username, Password: AuthUser.Password}
	configureJson, err := json.Marshal(configuration)
	logJsonMarshalError(err, jsonMarshalError)
	writeFile(ConfFileName, configureJson)
	log.Printf(connDetailsSuccessInfo, ConfFileName)
}

func getConnectionDetails() m.ConnDetails {
	var conf m.ConnDetails
	data := readFile(ConfFileName)
	err := json.Unmarshal([]byte(data), &conf)
	logJsonUnmarshalError(err, jsonUnmarshalError)
	return conf
}

func SetConnectionDetails() {
	if fileExists(ConfFileName) {
		conf := getConnectionDetails()
		NexusURL = conf.NexusURL
		AuthUser.Username = conf.Username
		AuthUser.Password = conf.Password
	} else {
		log.Printf(connDetailsEmptyInfo, "nexus3-repository-cli configure")
		os.Exit(1)
	}
}

func Usage() {
	flag.Usage = func() {
		fmt.Printf("Usage: nexus3-repository-cli [command]\n\n")
		fmt.Printf("[commands]\n  %s\t"+
			"%s\n  %s  \t"+
			"%s\n  %s  \t"+
			"%s\n  %s  \t"+
			"%s\n  %s  \t"+
			"%s\n  %s  \t%s"+
			"\n\n",
			ConfCommandFlag, ConfCommandUsage,
			ScriptCommandFlag, ScriptCommandUsage,
			RepoCommandFlag, RepoCommandUsage,
			SelectorCommandFlag, SelectorCommandUsage,
			PrivilegeCommandFlag, PrivilegeCommandUsage,
			RoleCommandFlag, RoleCommandUsage)
	}
}

func PrintConfCommandUsage(fs *flag.FlagSet) {
	fs.Usage = func() {
		fmt.Printf("Usage: nexus3-repository-cli configure [args]\n\n")
		fmt.Printf("[args]\n\n")
		fs.PrintDefaults()
		fmt.Printf("\n")
	}
}

func PrintScriptCommandUsage(fs *flag.FlagSet) {
	fs.Usage = func() {
		fmt.Printf("Usage: nexus3-repository-cli script [args] [options]\n\n")
		fmt.Printf("[args]\n\n  "+
			"-%s string\t%s\n  "+
			"-%s string\n\t%s\n  -%s string\n\t%s\n"+
			"\n[options]\n\n  -%s\n\t%s\n  -%s\n\t%s\n  -%s\n\t%s\n\n",
			TaskFlag, ScriptTaskUsage,
			ScriptNameFlag, ScriptNameUsage,
			ScriptPayloadFlag, ScriptPayloadUsage,
			SkipTlsFlag, SkipTlsUsage,
			DebugFlag, DebugUsage,
			VerboseFlag, VerboseUsage)
	}
}

func PrintRepoCommandUsage(fs *flag.FlagSet) {
	fs.Usage = func() {
		fmt.Printf("Usage: nexus3-repository-cli repo [args] [options]\n\n")
		fmt.Printf("[args]\n\n  "+
			"-%s string\t%s\n  "+
			"-%s string\n\t%s\n  -%s string\n\t%s\n  "+
			"-%s string\n\t%s\n  -%s string\n\t%s\n  "+
			"-%s string\n\t%s\n  -%s string\n\t%s\n  "+
			"-%s string\n\t%s\n  -%s string\n\t%s\n  "+
			"-%s string\n\t%s\n  -%s string\n\t%s\n  "+
			"\n[options]\n\n  -%s\n\t%s\n  -%s\n\t%s\n  -%s\n\t%s\n\n",
			TaskFlag, RepoTaskUsage,
			RepoNameFlag, RepoNameUsage,
			RepoFormatFlag, fmt.Sprintf(RepoFormatUsage, RepoFormats),
			RemoteURLFlag, RemoteURLUsage,
			RepoMembersFlag, RepoMembersUsage,
			ProxyUserFlag, ProxyUserUsage,
			ProxyPassFlag, ProxyPassUsage,
			DockerHttpPortFlag, DockerHttpPortUsage,
			DockerHttpsPortFlag, DockerHttpsPortUsage,
			BlobStoreNameFlag, BlobStoreNameUsage,
			ReleaseFlag, ReleaseUsage,
			SkipTlsFlag, SkipTlsUsage,
			DebugFlag, DebugUsage,
			VerboseFlag, VerboseUsage)
	}
}

func PrintSelectorCommandUsage(fs *flag.FlagSet) {
	fs.Usage = func() {
		fmt.Printf("Usage: nexus3-repository-cli selector [args] [options]\n\n")
		fmt.Printf("[args]\n\n  "+
			"-%s string\t%s\n  "+
			"-%s string\n\t%s\n  -%s string\n\t%s\n  -%s string\n\t%s\n"+
			"\n[options]\n\n  -%s\n\t%s\n  -%s\n\t%s\n  -%s\n\t%s\n\n",
			TaskFlag, SelectorTaskUsage,
			SelectorNameFlag, SelectorNameUsage,
			SelectorDescFlag, SelectorDescUsage,
			SelectorExpressionFlag, SelectorExpressionUsage,
			SkipTlsFlag, SkipTlsUsage,
			DebugFlag, DebugUsage,
			VerboseFlag, VerboseUsage)
	}
}

func PrintPrivilegeCommandUsage(fs *flag.FlagSet) {
	fs.Usage = func() {
		fmt.Printf("Usage: nexus3-repository-cli privilege [args] [options]\n\n")
		fmt.Printf("[args]\n\n  "+
			"-%s string\t%s\n  -%s string\n\t%s\n  "+
			"-%s string\n\t%s\n  -%s string\n\t%s\n  "+
			"-%s string\n\t%s\n  -%s string\n\t%s\n"+
			"\n[options]\n\n  -%s\n\t%s\n  -%s\n\t%s\n  -%s\n\t%s\n\n",
			TaskFlag, PrivilegeTaskUsage,
			PrivilegeNameFlag, PrivilegeNameUsage,
			PrivilegeDescFlag, PrivilegeDescUsage,
			PSelectorNameFlag, SelectorNameUsage,
			PRepoNameFlag, RepoNameUsage,
			ActionFlag, fmt.Sprintf(ActionUsage, PrivilegeActions),
			SkipTlsFlag, SkipTlsUsage,
			DebugFlag, DebugUsage,
			VerboseFlag, VerboseUsage)
	}
}

func PrintRoleCommandUsage(fs *flag.FlagSet) {
	fs.Usage = func() {
		fmt.Printf("Usage: nexus3-repository-cli role [args] [options]\n\n")
		fmt.Printf("[args]\n\n  "+
			"-%s string\t%s\n  -%s string\n\t%s\n  "+
			"-%s string\n\t%s\n  -%s string\n\t%s\n  "+
			"-%s string\n\t%s\n  -%s string\n\t%s\n"+
			"\n[options]\n\n  -%s\n\t%s\n  -%s\n\t%s\n  -%s\n\t%s\n\n",
			TaskFlag, RoleTaskUsage,
			RoleIDFlag, RoleIDUsage,
			RoleDescFlag, RoleDescUsage,
			RoleMembersFlag, RoleMembersUsage,
			RolePrivilegesFlag, RolePrivilegesUsage,
			UpdateActionFlag, fmt.Sprintf(UpdateActionUsage, UpdateActions),
			SkipTlsFlag, SkipTlsUsage,
			DebugFlag, DebugUsage,
			VerboseFlag, VerboseUsage)
	}
}
