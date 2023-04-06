SHELL=/bin/bash

# run this after changing proto files
# not working rn so just c/p
proto:
	export PATH="$PATH:$(go env GOPATH)/bin" && protoc -Iproto --go_out=. --go_opt=module=github.com/kainn9/grpc_game --go-grpc_out=. --go-grpc_opt=module=github.com/kainn9/grpc_game proto/players.proto

# Note: should probably use ->
# https://github.com/kainn9/go_grpc_multiplayer_game_example_attempt/actions/workflows/build_and_deploy_server.yaml
# *
# manual build with intent to be run on Linux ENV(like EC2)
# change GOOS otherwise
# will need 'chmod +x application' in ec2 
buildS:
	GOOS=linux GOARCH=amd64 go build -o bin/application ./server 

# Build client
buildC:
	go build -ldflags "-X github.com/kainn9/grpc_game/client_util.BuildTime=true" -o bin/application ./client && cp -R ./client/sprites ./bin && cp -R ./client/backgrounds ./bin && cp -R ./client/audio ./bin && chmod +x ./bin/application

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

# Note: should probably use ->
# https://github.com/kainn9/go_grpc_multiplayer_game_example_attempt/actions/workflows/build_and_deploy_server.yaml
# *
# manual build with intent to be run on Linux ENV(like EC2)
# change GOOS otherwise
# will need 'chmod +x application' in ec2 
buildSW:
	SET GOOS=linux& SET GOARCH=amd64& go build -o bin\application .\server\

# for windows
# ebitten does not support cross platform building(https://github.com/hajimehoshi/ebiten/discussions/1694#discussioncomment-964212)
# use github-action to get unix bin if needed:
# *
# https://github.com/kainn9/go_grpc_multiplayer_game_example_attempt/actions/workflows/build_and_upload_client.yaml
# *
# Reminder: github-action this only builds the bin file, assets still must be added into folder with application.bin(audio, sprites, backgrounds)
# you can swap the unix bin with application.exe after running buildCW, to acheive this more easily
buildCW:
	go build -ldflags "-X github.com/kainn9/grpc_game/client_util.BuildTime=true" -o bin\application.exe .\client\ && xcopy /E /I .\client\sprites bin\sprites\ && xcopy /E /I .\client\backgrounds bin\backgrounds\ && xcopy /E /I .\client\audio bin\audio\ && icacls bin\application.exe /grant:r "Users:(OI)(CI)F" /T

