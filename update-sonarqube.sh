#!/bin/env bash
go test ./... -short -coverprofile=coverage.out `go list ./.. | grep -v vendor/`
# sonar-scanner -Dsonar.login=ed8b8793d1b7c330cbef474b11cef082129416e8
