#!/bin/bash
go test api-server/apierror -v -count 1 -run "TestHandler"
