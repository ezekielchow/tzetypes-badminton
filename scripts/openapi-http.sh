#!/bin/bash
set -e

readonly service="$1"
readonly output_dir="$2"
readonly package="$3"

# Ensure output directory exists
mkdir -p "$output_dir"

# Run oapi-codegen in Docker
docker run --rm -v "$(pwd)":/app oapi-codegen-docker -generate types -o "/app/$output_dir/openapi_types.gen.go" -package "$package" "/app/api/$service.yml"
docker run --rm -v "$(pwd)":/app oapi-codegen-docker -generate chi-server,strict-server -o "/app/$output_dir/openapi_api.gen.go" -package "$package" "/app/api/$service.yml"
docker run --rm -v "$(pwd)":/app oapi-codegen-docker -generate client -o "/app/$output_dir/openapi_client_gen.go" -package "$package" "/app/api/$service.yml"
