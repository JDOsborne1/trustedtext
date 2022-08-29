docker rm $(docker stop $(docker ps -a -q --filter ancestor=tt_test))

docker run -d -p 8081:8080 --name first_env tt_test
docker run -d -p 8082:8080 --name second_env tt_test 