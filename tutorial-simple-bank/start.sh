#!/bin/sh

set -e

echo "Waiting for the database to be ready..."
/app/wait-for.sh db:5432 -t 300

echo "Running database migration..."
migrate -path ./db/migration -database "$DB_URL" -verbose up

echo "Starting the application..."
exec /app/main
