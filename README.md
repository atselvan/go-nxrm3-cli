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
  configure	Set nexus connection details
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

  -task string	Script Task (Required)  (For all tasks the script(s) should exist under the path ./scripts/groovy)

    list 	    List all the scripts available in Nexus. script-name (Optional) If script-name is passed the contents of the script will be printed
    add  	    Add a new script to nexus. script-name (Required)
    update 	    Update a script that is available in nexus. script-name (Required)
    add-or-update   Add or Update a script in nexus. script-name (Required)
    delete          Delete a script from nexus. script-name (Required)
    run 	    Run/Execute a script in nexus. Required Parameter: script-name

  -script-name string
	Name of the script to be executed in nexus. The script should exist under the path ./scripts/groovy
  -payload string
	Arguments to be passed to a nexus script can be sent as a payload during script execution

[options]

  -skip-tls
	Skip TLS verification for the nexus server instance
  -debug
	Set Default for more information on the nexus script execution
  -verbose
	Set Verbose for detailed http request and response logs
```

### repo sub-command

```console
Usage: nexus3-repository-cli repo [args] [options]

[args]

  -task string	Repo Task (Required)

    list   		List all the repositories in nexus.
			(Optional - repo-name) If repo-name is passed the list command will get the details of the repository.
			(Optional - repo-format) If repo-format is passed the list command will list the repositories as per the format
    create-hosted	Create a hosted repository in nexus. (Required - repo-name and repo-format)
    create-proxy	Create a proxy repository in nexus. (Required - repo-name, repo-format and remote-url ) (Optional - proxy-user and proxy-pass)
    create-group	Create a group repository in nexus. (Required - repo-name,repo-format and repo-members)
    add-group-members	Add new members to a existing group repository. (Required - repo-name,repo-format and repo-members)
    delete		Delete a repository from nexus

    If you are creating a docker repository it is necessary to also provide either a docker-http-port or a docker-https-port or both.

  -repo-name string
	Nexus repository name
  -repo-format string
	Repository format. Available formats : ["maven" "npm" "nuget" "bower" "pypi" "raw" "rubygems" "yum" "docker"]
  -remote-url string
	Remote URL to be proxied in nexus
  -repo-members string
	Comma-separated repository names that should be added to a group repo
  -proxy-user string
	Username for accessing the proxy repository
  -proxy-pass string
	Password for accessing the proxy repository
  -docker-http-port string
	Docker HTTP port
  -docker-https-port string
	Docker HTTPs port
  -blob-store-name string
	Blob store name
  -releases string
	Set this flag to create a releases repository
  
[options]

  -skip-tls
	Skip TLS verification for the nexus server instance
  -debug
	Set Default for more information on the nexus script execution
  -verbose
	Set Verbose for detailed http request and response logs

```