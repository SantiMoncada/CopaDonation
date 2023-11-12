# Stage 1: Build Go application
FROM golang:1.21.3 AS go_builder

WORKDIR /app

COPY go.mod /app

COPY go.sum /app

RUN go mod download

COPY . /app

RUN go build -o ServerExecutable src/*.go

# Stage 2: Build Node.js application
FROM node:18.18.2 AS node_builder

WORKDIR /app

COPY package.json /app

COPY package-lock.json /app

RUN npm install

COPY . /app

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
EXPOSE 3000
# Define the startup command for your application
CMD ["./ServerExecutable", "release"]
