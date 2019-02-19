#!/bin/bash

CLI="./nexus3-repository-cli"
scripts=("get-repo"
         "create-hosted-repo"
         "create-proxy-repo"
         "create-group-repo"
         "update-group-members"
         "delete-repo"
         "get-content-selectors"
         "create-content-selector"
         "update-content-selector"
         "delete-content-selector"
         "get-privileges"
         "create-privilege"
         "update-privilege"
         "delete-privilege"
         "get-roles"
         "create-role"
         "delete-role"
        )

for script in "${scripts[@]}"
do
    ${CLI} script -skip-tls -task add-or-update -name ${script}
done