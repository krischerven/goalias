#!/bin/bash
./vet.sh
go test ./...
go fmt ./...
