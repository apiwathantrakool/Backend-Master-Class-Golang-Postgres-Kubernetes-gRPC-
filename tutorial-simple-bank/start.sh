#!/bin/sh

set -e

echo "run db migration"
/app/wait-for.sh db:5432 -t 300
make migrateup "$DB_URL"
make server

echo "start the app"
exec "$@"