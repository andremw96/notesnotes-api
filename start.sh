#!/bin/sh

# script will exit immediately if comman return 0 status
set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@" #takes all parameter, pass to the script and run it