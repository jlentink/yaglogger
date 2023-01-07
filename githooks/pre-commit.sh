#!/bin/bash

STAGED_GO_FILES=$(git diff --cached --name-only -- '*.go')
if [[ $STAGED_GO_FILES == "" ]]; then
    echo "No Go Files to Update"
else
    for file in $STAGED_GO_FILES; do
        ## format our file
        go fmt "$file"
        ## add any potential changes from our formatting to the
        ## commit
        git add "$file"
    done
    golangci-lint run ./...
fi