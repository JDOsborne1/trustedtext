docker build -t tt_test .

docker run -d -p 8081:8080 tt_test
# Not needed for now
# docker run -d -p 8082:8080 tt_test 