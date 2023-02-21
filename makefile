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
	go build -ldflags "-X github.com/kainn9/grpc_game/util.BuildTime=true" -o bin/application ./client && cp -R ./client/sprites ./bin && cp -R ./client/backgrounds ./bin && cp -R ./client/audio ./bin && chmod +x ./bin/application

# run client
runC:
	cd ./client && 	go run .

runCR:
	cd ./client && 	go run . -race

# run server
runS:
	cd ./server && 	go run . 

runSR:
	cd ./server && 	go run . -race

genSSL:
	cd ./ssl && chmod +x ssl.sh && ./ssl.sh
	