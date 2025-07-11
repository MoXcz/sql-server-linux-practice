#!/bin/bash

echo "Waiting for SQL Server..."
./wait-for-it.sh bankDB:1433 --timeout=30 --strict -- echo "Server SQL is up."

echo "Running migrations..."
goose down
goose up

echo "Starting application..."
exec /usr/local/bin/app
