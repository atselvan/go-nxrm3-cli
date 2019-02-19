#!/bin/bash

CLI="./nexus3-repository-cli"
repoFormat=("maven" "npm" "nuget" "bower" "pypi" "raw" "rubygems" "yum")

cleanUpInitialRepositories() {
    # Initial repo cleanup
    $1 repo -skip-tls -task delete -name maven-central
    $1 repo -skip-tls -task delete -name maven-public
    $1 repo -skip-tls -task delete -name maven-releases
    $1 repo -skip-tls -task delete -name maven-snapshots
    $1 repo -skip-tls -task delete -name nuget-group
    $1 repo -skip-tls -task delete -name nuget-hosted
    $1 repo -skip-tls -task delete -name nuget.org-proxy
}

cleanUpRepoStructure(){
    $1 repo -skip-tls -task delete -name $2-snapshots
    $1 repo -skip-tls -task delete -name $2-releases
    $1 repo -skip-tls -task delete -name $2-proxy
    $1 repo -skip-tls -task delete -name $2-proxy-withCred
    $1 repo -skip-tls -task delete -name $2-group
}

cleanupDockerRepoStructure(){
    $1 repo -skip-tls -task delete -name docker-both
    $1 repo -skip-tls -task delete -name docker-http
    $1 repo -skip-tls -task delete -name docker-https
    $1 repo -skip-tls -task delete -name docker-proxy
    $1 repo -skip-tls -task delete -name docker-proxy-withCred
    $1 repo -skip-tls -task delete -name docker-group
}

cleanUpRoles(){
    $1 role -skip-tls -task delete -id maven-role
    $1 role -skip-tls -task delete -id npm-role
    $1 role -skip-tls -task delete -id nuget-role
    $1 role -skip-tls -task delete -id docker-role
}

cleanUpPrivileges(){
    $1 privilege -skip-tls -task delete -name maven-priv
    $1 privilege -skip-tls -task delete -name npm-priv
    $1 privilege -skip-tls -task delete -name nuget-priv
    $1 privilege -skip-tls -task delete -name docker-priv
}

cleanUpSelector(){
    $1 selector -skip-tls -task delete -name maven-selector
    $1 selector -skip-tls -task delete -name npm-selector
    $1 selector -skip-tls -task delete -name nuget-selector
    $1 selector -skip-tls -task delete -name docker-selector
}

#Cleanup
printf "****************************************************************************************************"
printf "\nDeleting Roles in Nexus\n"
printf "****************************************************************************************************\n"
cleanUpRoles ${CLI}

printf "****************************************************************************************************"
printf "\nDeleting Privileges\n"
printf "****************************************************************************************************\n"
cleanUpPrivileges ${CLI}

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
printf "****************************************************************************************************\n"