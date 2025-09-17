#!/bin/sh
set -e

echo "Starting application entrypoint..."

if [ -f .env ]; then
    echo " Loading environment variables from .env file"
    export $(cat .env | grep -v '^#' | xargs)
fi


DB_HOST=${DB_HOST:-postgres12}
DB_PORT=${DB_PORT:-5432}
DB_NAME=${DB_NAME:-vet_database}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-root}
DB_URL="postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

echo " Database Configuration:"
echo "   Host: ${DB_HOST}"
echo "   Port: ${DB_PORT}"
echo "   Database: ${DB_NAME}"
echo "   User: ${DB_USER}"

# Run migrations if migrate command is available
if command -v migrate &> /dev/null; then
    echo "Running database migrations..."
    
    echo "Waiting for PostgreSQL to be ready..."
    until pg_isready -h $DB_HOST -p $DB_PORT -U $DB_USER; do
        echo "Waiting for database connection..."
        sleep 2
    done
    
    echo " PostgreSQL is ready!"
    
    # Run migrations
    echo " Executing migrations..."
    migrate -path ./db/migrations -database "$DB_URL" -verbose up
    
    if [ $? -eq 0 ]; then
        echo "Migrations completed successfully!"
    else
        echo "Migrations failed with exit code: $?"
        echo " Continuing application startup..."
    fi
else
    echo " migrate command not found, skipping migrations"
fi

echo "Starting application..."
exec "$@"