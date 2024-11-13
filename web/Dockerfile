# Use an official Node.js image
FROM node:18 AS base
WORKDIR /app

# Copy package.json and yarn.lock
COPY package.json ./
# Install dependencies using Yarn
RUN yarn install --frozen-lockfile

# Copy the rest of the application code
COPY . .

# Use a non-root user for security
USER node

# Expose the port
EXPOSE 5173

# Start the development server
CMD ["yarn", "run", "dev", "--host"]
