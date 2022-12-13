#!/bin/bash

az group list | jq '.[].name' | xargs -I {} az group delete --resource-group {} --yes
rm -rf _output/ translations/
