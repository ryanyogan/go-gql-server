#!/bin/sh

srcPath="cmd"
pkgFile="main.go"
app="gql-server"
src="$srcPath/$app/$pkgFile"

printf "\nStar running: $app\n"

export $(grep -v '^#' .env | xargs) && time go run $src

unset $(grep -v '^#' .env | sed -E 's/(.*)=.*/\1/' | xargs)

printf "\nStopped Running: $app\n\n"
