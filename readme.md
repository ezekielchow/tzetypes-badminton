# Developer Setup
1. Install make
- On Ubuntu, you can install build essential tools which will also take care of installing the make:
```
sudo apt install build-essential
```
2. Install [docker](https://docs.docker.com/engine/install/) 
3. Install [docker compose](https://docs.docker.com/compose/install/)
4. Install [go](https://go.dev/doc/install). Version 1.22.5 or higher 
5. Run `docker compose up` in the root directory to start the services

# Directories & Files
- `api`: OpenAPI schema definitions
- `docker`: Dockerfiles for starting services in local development
- `gcp-proxy`: Proxy used to get identity token in live site to access api server
- `internal`: API server source files
- `scripts`: Helper scripts used by Makefile
- `web`: Frontend source files
- `Makefile`: Helper commands for project

# Making a pull request
1. Write & run test for any services for APIs
2. Run `make pre-commit` without any errors
3. Make PR!