#!/usr/bin/env bash

go run cmd/db-migrate/main.go --do create --db test_gorm
go run cmd/db-migrate/main.go --do migrate --db test_gorm