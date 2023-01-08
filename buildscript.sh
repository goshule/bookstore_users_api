
set GOARCH=amd64
set GOOS=linux
#set buildvcs=false
go build  -buildvcs=false -o bookstore-user-api .
#cp bookstore-user-api envshipp/
#docker build --no-cache -t sospetermule/bookstore-user-api:v002 .
#docker push sospetermule/bookstore-user-api:v002
#docker-compose -f  docker-compose-userapi.yml up -d
