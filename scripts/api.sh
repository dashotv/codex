#!/usr/bin/env bash

curl -s -H 'Content-type: application/json' \
  -H 'Accept: application/json' \
  "localhost:9020/api/$1" \
  -d "$2" | jq "${3:-.}"
