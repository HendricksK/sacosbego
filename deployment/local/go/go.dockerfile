FROM golang:latest

RUN mkdir $GOPATH/go-web

RUN sudo service goweb start

RUN sudo service goweb status 
