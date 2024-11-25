docker run --name mysql -d \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=admin \
  -e MYSQL_DATABASE=users \
  -e MYSQL_USER=abhishek \
  -e MYSQL_PASSWORD=abhishek \
  mysql:latest

# Number of times to send the request
N=10

# Loop to send the requests
for ((i=1; i<=N; i++)); do
    curl -s "http://localhost:8080/ping" &  # Run the curl command in the background 
done

wait 

docker rm -f mysql