#!/bin/bash

# Check if .env file exists
if [ ! -f .env ]; then
    echo "Error: .env file not found"
    exit 1
fi

# Read variables from .env file
source .env

atlas migrate diff $1 \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url $DATABASE_DEV_URL
