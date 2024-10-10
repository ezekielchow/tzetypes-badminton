#!/bin/bash
set -e

readonly service="$1"
readonly output_dir="$2"
readonly package="$3"

oapi-codegen -generate types -o "$output_dir/openapi_types.gen.go" -package "$package" "api/$service.yml"
oapi-codegen -generate chi-server,strict-server -o "$output_dir/openapi_api.gen.go" -package "$package" "api/$service.yml"
oapi-codegen -generate client -o "$output_dir/openapi_client_gen.go" -package "$package" "api/$service.yml"
