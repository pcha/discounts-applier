#!/bin/bash
go test ./...
docker-compose -f ./docker-compose-test.yml build
sh ./run-integration-tests.sh
docker-compose -f ./docker-compose-test.yml stop