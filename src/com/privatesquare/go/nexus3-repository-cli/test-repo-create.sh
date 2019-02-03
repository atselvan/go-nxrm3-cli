#!/bin/bash

CLI="./nexus3-repository-cli"
repoFormat=("maven" "npm" "nuget" "bower" "pypi" "raw" "rubygems" "yum")

createRepoStructure() {
    $1 repo -skip-tls -task create-hosted -repo-name $2-snapshots -repo-format $2
    $1 repo -skip-tls -task create-hosted -repo-name $2-releases -releases -repo-format $2
    $1 repo -skip-tls -task create-proxy -repo-name $2-proxy -releases -repo-format $2 -remote-url https://localhost
    $1 repo -skip-tls -task create-proxy -repo-name $2-proxy-withCred -releases -repo-format $2 -remote-url https://localhost -proxy-user test -proxy-pass test123
    $1 repo -skip-tls -task create-group -repo-name $2-group -repo-format $2 -repo-members $2-snapshots,$2-releases,$2-proxy,$2-proxy-withCred
}

cleanUpRepoStructure(){
    $1 repo -skip-tls -task delete -repo-name $2-snapshots
    $1 repo -skip-tls -task delete -repo-name $2-releases
    $1 repo -skip-tls -task delete -repo-name $2-proxy
    $1 repo -skip-tls -task delete -repo-name $2-proxy-withCred
    $1 repo -skip-tls -task delete -repo-name $2-group
}

createDockerRepoStructure(){
    $1 repo -skip-tls -task create-hosted -repo-name docker-http -repo-format docker
    $1 repo -skip-tls -task create-hosted -repo-name docker-both -repo-format docker -docker-http-port 18081 -docker-https-port 18444
    $1 repo -skip-tls -task create-hosted -repo-name docker-http -repo-format docker -docker-http-port 18080
    $1 repo -skip-tls -task create-hosted -repo-name docker-https -repo-format docker -docker-https-port 18443
    $1 repo -skip-tls -task create-proxy -repo-name docker-proxy -repo-format docker -docker-https-port 18445 -remote-url https://registry-1.docker.io
    $1 repo -skip-tls -task create-proxy -repo-name docker-proxy-withCred -repo-format docker -docker-https-port 18446 -remote-url https://registry-1.docker.io -proxy-user test -proxy-pass test123
    $1 repo -skip-tls -task create-group -repo-name docker-group -repo-format docker -docker-https-port 18447 -repo-members docker-both,docker-http,docker-https,docker-proxy,docker-proxy-withCred
}

cleanupDockerRepoStructure(){
    $1 repo -skip-tls -task delete -repo-name docker-both
    $1 repo -skip-tls -task delete -repo-name docker-http
    $1 repo -skip-tls -task delete -repo-name docker-https
    $1 repo -skip-tls -task delete -repo-name docker-proxy
    $1 repo -skip-tls -task delete -repo-name docker-proxy-withCred
    $1 repo -skip-tls -task delete -repo-name docker-group
}

cleanUpInitialRepositories() {
    # Initial repo cleanup
    $1 repo -skip-tls -task delete -repo-name maven-central
    $1 repo -skip-tls -task delete -repo-name maven-public
    $1 repo -skip-tls -task delete -repo-name maven-releases
    $1 repo -skip-tls -task delete -repo-name maven-snapshots
    $1 repo -skip-tls -task delete -repo-name nuget-group
    $1 repo -skip-tls -task delete -repo-name nuget-hosted
    $1 repo -skip-tls -task delete -repo-name nuget.org-proxy
}

#Cleanup
printf "****************************************************************************************************"
printf "\nDeleting initial Repositories in Nexus\n"
printf "****************************************************************************************************\n"
cleanUpInitialRepositories ${CLI}

printf "****************************************************************************************************"
printf "\nDeleting repository structures if created earlier\n"
printf "****************************************************************************************************\n"
for format in "${repoFormat[@]}"
do
    cleanUpRepoStructure ${CLI} ${format}
done
cleanupDockerRepoStructure ${CLI}

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