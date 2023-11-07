# Stage 1: Build Go application
FROM golang:1.21.3 AS go_builder

WORKDIR /app
COPY . /app

RUN go mod tidy

RUN go build -o ServerExecutable src/*.go

# Stage 2: Build Node.js application
FROM node:18.18.2 AS node_builder

WORKDIR /app

COPY . /app

RUN npm install

RUN npm run webpack:build
# Stage 3: Create a final image
FROM alpine

RUN apk add --no-cache libc6-compat 

COPY ./public ./app/public

COPY ./templates ./app/templates

# Copy the Go application from the first stage
COPY --from=go_builder /app/ServerExecutable /app

# Copy the Node.js application from the second stage
COPY --from=node_builder /app/public/dist /app/public/dist

# Set the working directory for the final image
WORKDIR /app

# Expose any necessary ports
EXPOSE 8080
# Define the startup command for your application
CMD ["./ServerExecutable", "release"]
