# Nexus3 Repository CLI

[![Build Status](https://travis-ci.org/atselvan/go-nxrm3-cli.svg?branch=master)](https://travis-ci.org/atselvan/go-nxrm3-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/atselvan/go-nxrm3-cli)](https://goreportcard.com/report/github.com/atselvan/go-nxrm3-cli)
[![Quality Gate](https://sonarcloud.io/api/project_badges/measure?project=io.sonarcloud.examples.go-sqscanner-travis-project&metric=alert_status)](https://sonarcloud.io/dashboard/index/atselvan:go-nxrm3-cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A easy to use CLI written in golang which is intended to automate all the processes around Sonatype Nexus 3 Repository 
Manager. The CLI is built using 

## Required Libraries

[atselvan/go-nxrm-lib](https://github.com/atselvan/go-nxrm-lib)

Package go-nxrm-lib implements functions to call Nexus repository manager 3 and provision resources in nexus using the Integration API of nexus (scripts API)

[spf13/cobra](https://github.com/spf13/cobra)

Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.

## Features

* Add, List, Update, Delete Scripts
* Create, List and Delete Hosted, Proxy and group repositories in Nexus
* Add or Remove members to/from an existing group repository
* Create, List, Update and Delete Content selectors
* Create, List, Update and Delete repository content selector privileges
* Create, List, Update and Delete roles
* Add or Remove role members or privileges to/from an existing role

## Building the CLI

```bash
# From the root of the repository
export GOPATH=$(pwd)

# From the directory containing the main.go file
go get ./...
go build
```

For building the cli for multiple distributions checkout [mitchellh/gox](https://github.com/mitchellh/gox)

## Configuration

```cmd
./nexus3-repository-cli configure --nexus-url http://localhost:8081 --username admin --password admin

2019/08/12 21:43:28 Connection details were stored successfully in the file ./nexus3-repository-cli.json
```

The connection details are saved in a file parallel to the CLI. Once you are done with your commands make sure to 
delete the credentials.

## Examples

### Creating a repository

```cmd
./nexus3-repository-cli repo create --name Releases --type hosted --format maven --releases

2019/08/12 21:45:40 Repository "Releases" was created in nexus
```
### Listing roles

```cmd
./nexus3-repository-cli role list

nx-anonymous
nx-admin
2019/08/12 21:46:38 Number of roles in nexus : 2

./nexus3-repository-cli role list --id nx-admin

Role Details:
ID: nx-admin
Name: 
Description: Administrator Role
Source: default
Roles: []
Privileges: [nx-all]
```

## Help

```cmd
./nexus3-repository-cli -h

CLI to interacts with Nexus repository Manager 3
via its API to administer the instance and to create nxrm components

Usage:
  nexus3-repository-cli [command]

Available Commands:
  configure   Set nexus connection details
  scripts     Nexus script operations
  repo        Nexus repository operations
  selector    Nexus content selector operations
  privilege   Nexus privilege operations
  role        Nexus role operations
  help        Help about any command

Flags:
  -d, --debug      Set Default for more information on the nexus script execution
  -h, --help       help for nexus3-repository-cli
  -k, --skip-tls   Skip TLS verification for the nexus server instance
  -v, --verbose    Set Verbose for detailed http request and response logs

Use "nexus3-repository-cli [command] --help" for more information about a command.

```