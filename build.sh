#!/bin/bash

set -xeuf -o pipefail

go generate ./...
go test ./...
go build ./cmd/ethernaut
go build ./cmd/4word
