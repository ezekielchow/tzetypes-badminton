#!/bin/bash
set -e

readonly service="$1"

docker run --rm --env "JAVA_OPTS=-Dlog.level=error" -v "${PWD}:/local" \
  "openapitools/openapi-generator-cli:v7.8.0" generate \
  -i "/local/api/$service.yml" \
  -g typescript-fetch \
  -o "/local/web/src/repositories/clients/$service"
