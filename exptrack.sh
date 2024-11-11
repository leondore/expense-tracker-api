#!/bin/bash

case "$1" in
    start)
        echo "Starting Database container..."
        docker-compose up -d
        ;;
    stop)
        echo "Stopping Database container..."
        docker-compose down
        ;;
    migrate-up)
        echo "Running migrations..."
        goose -dir ./database/schema postgres "postgres://db_admin:password@localhost:6543/exptrack_db?sslmode=disable" up
        ;;
    migrate-down)
        echo "Reverting migrations..."
        goose -dir ./database/schema postgres "postgres://db_admin:password@localhost:6543/exptrack_db?sslmode=disable" down
        ;;
    migrate-create)
        echo "Creating migration..."
        goose -dir ./database/schema create "$2" sql
        ;;
    db-generate)
        echo "Generating database code..."
        sqlc generate
        ;;
    *)
        echo "Usage: $0 {start|stop|migrate-up|migrate-down}"
        exit 1
esac