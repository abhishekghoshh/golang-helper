#!/bin/bash

# Define container names
REDIS_CONTAINER="redis"
MONGODB_CONTAINER="mongodb"
POSTGRESQL_CONTAINER="postgresql"

# Function to stop and remove a container
remove_container() {
    local container_name=$1
    
    echo "Stopping container: $container_name..."
    docker stop $container_name
    if [ $? -ne 0 ]; then
        echo "Failed to stop container: $container_name. It may not be running."
    fi

    echo "Removing container: $container_name..."
    docker rm $container_name
    if [ $? -ne 0 ]; then
        echo "Failed to remove container: $container_name."
    else
        echo "Container $container_name removed successfully."
    fi
}

# Remove Redis container
remove_container $REDIS_CONTAINER

# Remove MongoDB container
remove_container $MONGODB_CONTAINER

# Remove PostgreSQL container
remove_container $POSTGRESQL_CONTAINER

echo "All specified containers have been processed."