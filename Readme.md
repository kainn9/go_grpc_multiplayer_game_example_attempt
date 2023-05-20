## Game Demo(5/7)
https://user-images.githubusercontent.com/85503587/236705026-3f053803-35cb-465e-9791-850127b56687.mp4

## Disclaimer(2/13/23 & beyond)

This Project could be a lot better. It's messy, and presumably, it's not following best practices, at least not yet. GRPC/HTTP2 is likely a weird, perhaps even bad, choice for a real-time multiplayer fighter game. I am not a game developer, and aside from tinkering around as a child, I have no experience/no idea what I'm doing here(I would like to be an indie game dev guy once I retire though haha).

At its core, this boilerplate/game is my way of embarking on the journey of learning Golang + the other stuff(but mostly Go) through the lens of an old passion. After all, my initial love for creating software stems from my love of video games as a child. MMORPGs from 2007-2010 have a special place in my heart. This Project will not be able to support an MMO scale(obviously lol), but perhaps some tiny multiplayer shenanigans.

## Why?
Professionally, I work primarily with TypeScript/Ruby for web development, and I have a lot of room to grow in that department too. I wanted a side project that could help me grow professionally while being fun on a personal level: Golang and gRPC's bidirectional streaming seemed like something cool worth exploring.

## Controls
### Player Controls
- Arrow keys for movement
- Space to jump/wall-jump
- Shift for defense(if applicable to character)
- Q, W, E, R for attacks

### Client Controls
- z/x to control volume
- M key to hide instructions(its a wall of text on the screen)

### Admin/Dev Controls
- N key to change character type
- U key to enable random world spawn
- I key to enable village spawn(random spawn setting takes precedence if also toggled)
- 4 key to swap world/arena
- 1 for dev mode camera
- 2 to rotate dev ruler(when dev move cam is active)
- w/s to increase/decrease ruler/camera speed(when dev move cam is active)
- 3 key to preview client side geometry placement(devWorldBuilder)
- L key for hitbox preview tool


## Local Dev
- In order to develop locally you must have Golang + Protoc installed
- I work on both unix and windows(so either should be okay), but the makefile is geared towards unix
- if you need to update the proto file, run 
```bash
protoc -Iproto --go_out=. --go_opt=module=github.com/kainn9/grpc_game --go-grpc_out=. --go-grpc_opt=module=github.com/kainn9/grpc_game proto/players.proto
```
- Otherwise to get started:
  - Install deps with
```bash
go mod download
go mod tidy
```
- run server with:
```bash
make runS
-or-
cd server
go run .
```
then run client with:
```bash
make runC
-or-
cd server
go run .
```

### Local "Dev Tools"(they're a bit janky)

#### Hitbox Tool
You can modify this [file](https://github.com/kainn9/go_grpc_multiplayer_game_example_attempt/blob/main/client/hitBoxTest.go) and use the L key to preview hitboxes(clientside). The code can then be used to create attack/hitbox sequences in the server. The file also contains some examples. Here is a video:

https://user-images.githubusercontent.com/85503587/227417024-60006b07-db70-47e8-9f08-d180fa982ac7.mp4

### World Builder Tool
In order to figure out where to place geometry on the backend, there is a frontend "dev cam" + ruler thing(toggled by the 1 key). You can use it to measure/find cords to place geometry. You can use the [devWorldBuilder function](https://github.com/kainn9/go_grpc_multiplayer_game_example_attempt/blob/main/client/devWorldBuilder.go#L11) to place/preview geometry on the clientside(press 3 key when dev cam is active to render previews). The calc is broken, so currently it only works if you haven't moved yet(you can use the 4 key to reset player position if you have moved). I'm not very proud of this tool, and would like to replace it with a tilemap system or something better in the future. Heres a video example:

https://user-images.githubusercontent.com/85503587/227417644-ac291714-92b1-4e58-bdef-9a32eb759f7b.mp4



## Library Credz
- https://github.com/hajimehoshi/ebiten <-- 2D GOLANG GAME Library
- https://github.com/SolarLune/resolv <- collison lib for GO(commonly used w/ ebiten). The backend platformer "physics" is heavily based on the repo's [world platformer example](https://github.com/SolarLune/resolv/blob/master/examples/worldPlatformer.go)

## Asset Credz
- Knight Sprite: https://szadiart.itch.io/2d-soulslike-character
- Monk Sprite: https://chierit.itch.io/elementals-ground-monk
- Demon Sprite: https://chierit.itch.io/boss-demon-slime
- Mage: https://chierit.itch.io/monthly-character-002
- Werewolf: https://chierit.itch.io/mc-003-werewolf
- Heavy Knight: https://luizmelo.itch.io/heavy-armor
- Bird Droid: https://penusbmic.itch.io/sci-fi-character-pack-5
- Action Bar Stuff: https://penusbmic.itch.io/the-dark-series-skill-icons-numbers
- Portals: https://creativekind.itch.io/animated
- Stun Effect Sprite: https://bdragon1727.itch.io/fire-pixel-bullet-16x16
- Asset Pack used to create worlds/arenas: https://itch.io/s/88510/ssworlds-bundle
- In Game Music: https://greenpiccolo.bandcamp.com/track/ride-with-me
- Game Demo Video Song: https://www.youtube.com/watch?v=QZluIr6PeEA
