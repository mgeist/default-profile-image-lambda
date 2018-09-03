#!/bin/bash

export GOOS=linux 
export GOARCH=amd64 
go build
rm default-profile-image-lambda.zip
zip -r default-profile-image-lambda.zip default-profile-image-lambda font.ttf
