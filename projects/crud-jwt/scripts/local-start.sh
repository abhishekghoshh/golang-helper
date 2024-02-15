sh scripts/local-destory.sh

SCRIPTS_DIR=$(dirname $0)
ENV_FILE_PATH=${PWD}/${SCRIPTS_DIR}/env

docker network create crud

docker run -d \
    -p 3307:3306 \
    --name localDB \
    --network crud \
    -e MYSQL_ROOT_PASSWORD=root \
    -e MYSQL_DATABASE=students \
    -e MYSQL_USER=abhishek \
    -e MYSQL_PASSWORD=abhishek \
    mysql:8.0

docker build -t crud .
sleep 10
docker run -d \
    -p 8090:8080 \
    --network crud \
    --name crud \
    --env-file=${ENV_FILE_PATH} \
    crud:latest

sleep 3

curl --location 'http://localhost:8090/account' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userName": "abhishekghosh",
    "password":"admin123@test",
    "name": "Abhishek Ghosh",
    "address": "18 No Alep Khan Mahalla Road",
    "phoneNo": "9563991640",
    "email": "ghoshabhishek1640@gmail.com",
    "dateOfBirth": "1997-09-10"
}'

curl --location 'http://localhost:8090/auth/login' \
--header 'Content-Type: application/json' \
--header 'Cookie: XSRF-TOKEN=092a9a17-797d-4b7c-ae49-6903605dbd59' \
--data '{
    "username":"YWJoaXNoZWtnaG9zaA==",
    "password":"YWRtaW4xMjNAdGVzdA=="
}'


docker logs -f crud