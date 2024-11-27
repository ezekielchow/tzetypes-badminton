# Use a lightweight Go image
FROM golang:1.22.5

# Install git and oapi-codegen
RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/* && \
    go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

# Set working directory
WORKDIR /app

# Entrypoint for the container
ENTRYPOINT ["oapi-codegen"]
