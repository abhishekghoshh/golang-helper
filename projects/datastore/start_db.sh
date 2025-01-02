#!/bin/bash

# Redis information
REDIS_CONTAINER="redis"
REDIS_PORT=6379

# MongoDB information
MONGODB_CONTAINER="mongodb"
MONGODB_PORT=27017

# postgres information
POSTGRES_CONTAINER="postgresql"
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres




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
docker run --name $POSTGRES_CONTAINER -p $POSTGRES_PORT:5432 \
    -e POSTGRES_USER=$POSTGRES_USER -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
    -e POSTGRES_DB=$POSTGRES_DB -d postgres
if [ $? -ne 0 ]; then
    echo "Failed to start PostgreSQL container."
    exit 1
fi

echo "All containers started successfully."
