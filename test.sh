#!/bin/bash
go mod tidy && go test -timeout 5s -run . go-montgomery-reduce -v -count=1