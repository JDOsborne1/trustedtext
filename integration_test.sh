docker compose up -d
cd trustedtext-webserver
go test 
cd ..
docker compose down