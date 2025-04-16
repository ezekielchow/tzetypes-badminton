# About
A badminton score tracking & post analysis app

# Demo
![Screenshot from 2025-04-16 10-29-10](https://github.com/user-attachments/assets/3bce7582-3fa0-4ed5-95f6-cc83bf234194)
![Screenshot from 2025-04-16 10-29-48](https://github.com/user-attachments/assets/ffab294e-5753-4906-b9d8-d1a473da7202)
![Screenshot from 2025-04-16 10-30-07](https://github.com/user-attachments/assets/325d46ed-c769-4d85-af32-5dfa1a9fda90)


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
