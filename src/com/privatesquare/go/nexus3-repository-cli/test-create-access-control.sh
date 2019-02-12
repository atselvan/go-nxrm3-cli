#!/usr/bin/env bash

CLI="./nexus3-repository-cli"

printf "****************************************************************************************************\n"
$CLI selector -skip-tls -task create -name maven-selector -expression "format == \"maven2\" and coordinate.groupId =^ \"com.company\""
$CLI selector -skip-tls -task create -name npm-selector -expression "format == \"npm\" and coordinate.groupId =^ \"com.company\""
$CLI selector -skip-tls -task create -name nuget-selector -expression "format == \"nuget\" and coordinate.groupId =^ \"Company\""
$CLI selector -skip-tls -task create -name docker-selector -expression "format == \"docker\" and coordinate.groupId =^ \"company/\""
printf "****************************************************************************************************\n"
printf "****************************************************************************************************\n"
$CLI privilege -skip-tls -task create -name maven-priv -selector-name maven-selector -repo-name maven-releases
$CLI privilege -skip-tls -task create -name npm-priv -selector-name npm-selector -repo-name npm-releases
$CLI privilege -skip-tls -task create -name nuget-priv -selector-name nuget-selector -repo-name nuget-releases
$CLI privilege -skip-tls -task create -name docker-priv -selector-name docker-selector -repo-name docker-both -action read
printf "****************************************************************************************************\n"
printf "****************************************************************************************************\n"
$CLI role -skip-tls -task create -id maven-role -role-privileges maven-priv
$CLI role -skip-tls -task create -id npm-role -role-privileges npm-priv
$CLI role -skip-tls -task create -id nuget-role -role-privileges nuget-priv
$CLI role -skip-tls -task create -id docker-role -role-privileges docker-priv
printf "****************************************************************************************************\n"
