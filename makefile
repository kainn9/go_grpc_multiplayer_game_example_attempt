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

# run client should work linux or windows
runC:
	cd ./client && 	go run .

runCR:
	cd ./client && 	go run . -race

# run server should work linux or windows
runS:
	cd ./server && 	go run . 

runSR:
	cd ./server && 	go run . -race

genSSL:
	cd ./ssl && chmod +x ssl.sh && ./ssl.sh

# Attempted windows versions of commands(might be broken)
protoW:
	protoc -Iproto --go_out=. --go_opt=module=github.com/kainn9/grpc_game --go-grpc_out=. --go-grpc_opt=module=github.com/kainn9/grpc_game proto/players.proto

buildSW:
	SET GOOS=linux& SET GOARCH=amd64& go build -o bin\application.exe .\server\

buildCW:
	go build -ldflags "-X github.com/kainn9/grpc_game/util.BuildTime=true" -o bin\application.exe .\client\ && xcopy /E /I .\client\sprites bin\sprites\ && xcopy /E /I .\client\backgrounds bin\backgrounds\ && xcopy /E /I .\client\audio bin\audio\ && icacls bin\application.exe /grant:r "Users:(OI)(CI)F" /T

