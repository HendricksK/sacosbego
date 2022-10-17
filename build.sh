#!/bin/bash
#make sure master is up to date
echo "git pull"
git pull
#getting env variables
echo "building creds into environmental variables"
source ./sacos-creds/make_creds.sh
#build package into main
echo "run build sacos golang"
go build -o main

