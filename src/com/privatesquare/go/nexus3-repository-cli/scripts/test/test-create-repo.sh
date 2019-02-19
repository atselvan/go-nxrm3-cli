#!/bin/bash

CLI="./nexus3-repository-cli"
repoFormat=("maven" "npm" "nuget" "bower" "pypi" "raw" "rubygems" "yum")

createRepoStructure() {
    $1 repo -skip-tls -task create-hosted -name $2-snapshots -format $2
    $1 repo -skip-tls -task create-hosted -name $2-releases -releases -format $2
    $1 repo -skip-tls -task create-proxy -name $2-proxy -releases -format $2 -remote-url https://localhost
    $1 repo -skip-tls -task create-proxy -name $2-proxy-withCred -releases -format $2 -remote-url https://localhost -proxy-user test -proxy-pass test123
    $1 repo -skip-tls -task create-group -name $2-group -format $2 -members $2-snapshots,$2-releases,$2-proxy-withCred
}

addGroupMembers(){
    $1 repo -skip-tls -task add-group-members -name $2-group -format $2 -members $2-snapshots,$2-releases,$2-proxy,$2-group
}

createDockerRepoStructure(){
    $1 repo -skip-tls -task create-hosted -name docker-http -format docker
    $1 repo -skip-tls -task create-hosted -name docker-both -format docker -docker-http-port 18081 -docker-https-port 18444
    $1 repo -skip-tls -task create-hosted -name docker-http -format docker -docker-http-port 18080
    $1 repo -skip-tls -task create-hosted -name docker-https -format docker -docker-https-port 18443
    $1 repo -skip-tls -task create-proxy -name docker-proxy -format docker -docker-https-port 18445 -remote-url https://registry-1.docker.io
    $1 repo -skip-tls -task create-proxy -name docker-proxy-withCred -format docker -docker-https-port 18446 -remote-url https://registry-1.docker.io -proxy-user test -proxy-pass test123
    $1 repo -skip-tls -task create-group -name docker-group -format docker -docker-https-port 18447 -members docker-both,docker-http,docker-https,docker-proxy,docker-proxy-withCred
}

printf "****************************************************************************************************"
printf "\nCreating repository structures\n"
printf "****************************************************************************************************\n"
for format in "${repoFormat[@]}"
do
    createRepoStructure ${CLI} ${format}
done

printf "****************************************************************************************************"
printf "\nCreating docker repository structures\n"
printf "****************************************************************************************************\n"
createDockerRepoStructure ${CLI}

printf "****************************************************************************************************"
printf "\nAdding Members to a group\n"
printf "****************************************************************************************************\n"
for format in "${repoFormat[@]}"
do
    addGroupMembers ${CLI} ${format}
done