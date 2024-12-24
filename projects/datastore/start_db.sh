#!/bin/bash

# Define container names
REDIS_CONTAINER="redis"
MONGODB_CONTAINER="mongodb"
POSTGRESQL_CONTAINER="postgresql"

# Define ports
REDIS_PORT=6379
MONGODB_PORT=27017
POSTGRESQL_PORT=5432

# Run Redis container
echo "Starting Redis container..."
docker run --name $REDIS_CONTAINER -p $REDIS_PORT:$REDIS_PORT -d redis
if [ $? -ne 0 ]; then
    echo "Failed to start Redis container."
    exit 1
fi

# Run MongoDB container
echo "Starting MongoDB container..."
docker run --name $MONGODB_CONTAINER -p $MONGODB_PORT:$MONGODB_PORT -d mongo
if [ $? -ne 0 ]; then
    echo "Failed to start MongoDB container."
    exit 1
fi

# Run PostgreSQL container
echo "Starting PostgreSQL container..."
docker run --name $POSTGRESQL_CONTAINER -p $POSTGRESQL_PORT:$POSTGRESQL_PORT \
    -e POSTGRES_USER=myuser -e POSTGRES_PASSWORD=mypassword \
    -e POSTGRES_DB=mydatabase -d postgres
if [ $? -ne 0 ]; then
    echo "Failed to start PostgreSQL container."
    exit 1
fi

echo "All containers started successfully."
