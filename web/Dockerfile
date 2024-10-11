# use the official Bun image
# see all versions at https://hub.docker.com/r/oven/bun/tags
FROM oven/bun:1 AS base
WORKDIR /app

COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

COPY . .

USER bun
EXPOSE 5173
ENTRYPOINT [ "bun", "run", "dev", " --host" ]