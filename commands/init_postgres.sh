#!/bin/bash

# This script is used to initialize the postgres database
psql -h hostname -U username -f "${PWD}"/internal/pkg/script/migrations.sql