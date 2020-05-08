#!/bin/bash
./vet.sh
go test ./...
./fmt.sh
