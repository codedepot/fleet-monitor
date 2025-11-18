#/bin/bash

rm -rf ./client

if [[ "$OSTYPE" == "msys" ]]; then
    export MSYS_NO_PATHCONV=1
fi
export GIT_USER_ID=codedepot
docker run --rm -v ./:/local openapitools/openapi-generator-cli generate \
    -i /local/openapi.json \
    -g go \
    -p packageName=client \
    -o /local/client/ \
    --git-repo-id fleet-monitor/client --git-user-id codedepot \
    
# do not make it its own module for ease of import
rm ./client/go.mod
rm ./client/go.sum