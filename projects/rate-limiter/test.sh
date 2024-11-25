#!/bin/bash

# Number of times to send the request
N=10  # You can change this value to the desired number of requests

# Loop to send the requests
for ((i=1; i<=N; i++)); do
    echo "Request #$i:"
    curl -s http://localhost:8080/ping/rate-limit-tollbooth  # The -s option silences progress output
    echo "\n"  # Adds a newline for better readability
done