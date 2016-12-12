#!/usr/bin/env bash
cd migration/
goose postgres "user=postgres dbname=postgres sslmode=disable" up
