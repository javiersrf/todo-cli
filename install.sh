#!/bin/bash

# Build the Go project
go build -o bin/todo-cli

# Move the binary to /usr/local/bin with sudo
sudo mv ./bin/todo-cli /usr/local/bin/