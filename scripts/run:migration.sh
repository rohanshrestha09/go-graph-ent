#!/bin/bash

# Check if .env file exists
if [ ! -f .env ]; then
    echo "Error: .env file not found"
    exit 1
fi

# Read variables from .env file
source .env

atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url $DATABASE_URL \
