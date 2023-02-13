### ReadeME(WIP)

## CREDZ

- https://github.com/hajimehoshi/ebiten <-- 2D GOLANG GAME ENGINE
- https://github.com/SolarLune/resolv <- collison lib for GO(commonly used w/ ebiten). The backend platformer "physics" is basically a rip from the world platformer exmaple in the repo

### ASSETS Creds

- https://szadiart.itch.io/2d-soulslike-character
- https://itch.io/s/88510/ssworlds-bundle

## Disclaimer(2/13/23)

This Project could be a lot better. It's messy, and presumably, it's not following best practices, at least not yet. I am not a game developer, and aside from tinkering around as a child, I have no experience/no idea what I'm doing here. Furthermore, I'm also a Golang/gRPC/Protobuf noob. At its core, this boilerplate/game is my way of embarking on the journey of learning Golang through the lens of an old passion. After all, my initial love for creating software stems from my love of video games as a child. MMORPGs from 2007-2010 have a special place in my heart. This Project will not be able to support an MMO scale, but perhaps some tiny multiplayer shenanigans if I get lucky.

## Why?

Professionally, I work primarily with TypeScript/Ruby for web development, and I have a lot of room to grow in that department too. I wanted a side project that could help me grow professionally while being fun on a personal level: Golang and gRPC's bidirectional streaming seemed like something worth exploring. If I manage to get this thing working somewhat reasonably, I'm hoping it can be a helpful boilerplate or reference for people interested in this kind of thing...maybe it can save someone a bit of time one day.

## Does It Work?

Ehhh, just barely. This thing is in its infancy and still in the POC stages.

### DEMO CLIENTS

- [MAC CLIENT DEMO DOWNLOAD](https://github.com/kainn9/go_grpc_multiplayer_game_example_attempt/files/10719870/macDemo.zip)
- [WINDOWS CLIENT DEMO DOWNLOAD](https://github.com/kainn9/go_grpc_multiplayer_game_example_attempt/files/10719872/windowsDemo.zip)

^ you should be able to unzip the files and double click on `application` to run the client. Hopefully my server is still up 🤞.

### Quick "GamePlay" Example(TODO)
https://user-images.githubusercontent.com/85503587/218402441-765416f9-a8a7-4ec8-a18f-684a840cd22d.mp4



### Quick World Building Example(its really bad lol)(TODO)
https://user-images.githubusercontent.com/85503587/218402479-4cc569e5-97a0-4ee4-a434-65801021cf6b.mp4

### General Concept:

- Client has bidirectional stream with server
- Clients streams inputs of a player with their randomly generated ID to server
- Server updates "state"(resolv space) based on players inputs
- Server streams back state to Client
- Client renders state from Server

### TODO: local environment setup

(if using unix based, the makefile should be enough to get started for now)

### TODO: Tests, Diagrams/Planning Docs, Goals/Roadmap(basically everything LOL).
