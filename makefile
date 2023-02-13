SHELL=/bin/bash

# run this after changing proto files
# not working rn so just c/p
proto:
	export PATH="$PATH:$(go env GOPATH)/bin" && protoc -Iproto --go_out=. --go_opt=module=github.com/kainn9/grpc_game --go-grpc_out=. --go-grpc_opt=module=github.com/kainn9/grpc_game proto/players.proto

# build server
buildS:
	GOOS=linux GOARCH=amd64 go build -o bin/application ./server 

# Build client
buildC:
	go build -ldflags "-X main.BuildTime=true" -o bin/application ./client && cp -R ./client/sprites ./bin && cp -R ./client/backgrounds ./bin && chmod +x ./bin/application

# run client
runC:
	cd ./client && 	go run .

# run server
runS:
	cd ./server && 	go run . 

genSSL:
	cd ./ssl && chmod +x ssl.sh && ./ssl.sh
	