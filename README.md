# Nexus3 Repository CLI

This CLI is intended to automate the processes around Sonatype Nexus 3 Repository Manager.

## Building the CLI

```console
export GOPATH=$(pwd) <From the root of the repository>
go build
```

## Using the CLI

```console
Usage: nexus3-repository-cli [command]

[commands]
  configurate	Set nexus connection details
  script  	Nexus script operations
  repo  	Nexus repository operations
```

### configure sub-command

```console
Usage: nexus3-repository-cli configure [args]

[args]
  -nexus-url string
    	Nexus 3 server URL. (Required)
  -password string
    	Nexus 3 server login password. (Required)
  -username string
    	Nexus 3 server login user. (Required)
```

**Example:**

```console
nexus3-repository-cli configure -nexus-url http://localhost:8081 -username admin -password admin123
```

### script sub-command

```console
Usage: nexus3-repository-cli script [args] [options]

[args]
  -task string
	Script Task (list|add|update|add-or-update|delete|run). (Required)
  -script-name string
	Name of the script to be executed in nexus. The script should exist under the path ./scripts/groovy.
  -payload string
	Arguments to be passed to a nexus script can be sent as a payload during script execution.

[options]
  -skip-tls
	Skip TLS verification for the nexus server instance.
  -debug
	Set Default for more information on the nexus script execution.
  -verbose
	Set Verbose for detailed http request and response logs.
```

### repo sub-command

```console
Usage: nexus3-repository-cli repo [args] [options]

[args]
  -repo-name string
	Nexus repository name.
  -repo-format string
	Repository format (maven|npm|nuget|docker).
  -remote-url string
	Remote URL to be proxied in nexus.
  -repo-members string
	Comma-separated repository names that should be added to a group repo.
  -release
	Set this flag to create a releases maven repository.
[options]
  -skip-tls
	Skip TLS verification for the nexus server instance.
  -debug
	Set Default for more information on the nexus script execution.
  -verbose
	Set Verbose for detailed http request and response logs.
```