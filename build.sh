#!/bin/bash
#make sure master is up to date
git pull
#build package into main
go build -o main

