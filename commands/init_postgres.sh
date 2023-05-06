#!/bin/bash

# This script is used to initialize the postgres database
psql -h localhost -d osg_arch -U postgres -W -f "${PWD}"/internal/pkg/script/migrations.sql