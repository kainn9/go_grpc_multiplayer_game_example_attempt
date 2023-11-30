## Game Demo(5/7)
https://user-images.githubusercontent.com/85503587/236705026-3f053803-35cb-465e-9791-850127b56687.mp4

## Please Read
While working on this project has been an incredibly enjoyable experience, I must confess that it is a bit of a chaotic endeavor. Whether you look at it from an architectural, design, or implementation perspective, it doesn't make sense. However, despite its clear messiness, I want to emphasize that this project has been far from a waste of time. Going into it, I had several specific goals in mind:

Get Comfortable with Golang: The project provided an excellent opportunity for me to become more proficient in Golang.

- Learn Protobuf/GRPC: Exploring Protobuf/GRPC was a significant aspect of my learning journey, even if it might not have been the most suitable choice for real-time multiplayer.

- Explore Gaming Netcode: The endeavor allowed me to delve into gaming netcode, discovering both its challenges and opportunities.

- Host a Multiplayer Gaming Server via AWS: Setting up and managing a multiplayer gaming server on AWS was a valuable experience that added a practical dimension to my skills.

Admittedly, using http2/GRPC for real-time multiplayer might not have been the most efficient choice, as protocols like UDP or a custom TCP setup could have offered less overhead. However, the project presented the chance to experiment with bidirectional GRPC streams, showcasing the versatility of the technology.

Yes, I might have organized packages and files without a deep understanding of proper game design paradigms initially. Nonetheless, this led me to delve into the Entity-Component-System (ECS) pattern, which has become a fundamental part of my new project/game and physics library.

Acknowledging that  I took advantage of Golang's concurrency features to create a complex web of time-based callbacks and race conditions, it did provide an opportunity to learn more about the sync package.

In essence, this project represents a journey of learning what to do by initially implementing what not to do—a common aspect of the learning process. If you decide to explore the code, don't expect perfection; it may not be pretty, elegant, or robust, and that's perfectly fine. Hidden beneath the layers of imperfection, there might be some valuable insights. For instance, I surprisingly find merit in the simple CI/CD GitHub actions, which stands out as a positive outcome amidst the challenges encountered during this development journey.



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
