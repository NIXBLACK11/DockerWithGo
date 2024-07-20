# DockerWithGo
Making linux containers with Golang

docker         run <container> cmd args
go run main.go run cmd args

/bin/bash -> creates a new running instance

command to run the container
```bash
sudo go run main.go run /bin/bash
```
build using this image -> [ubuntu-base-14.04](ubuntu-base-14.04-core-i386.tar.gz)

Reference -> [Liz Rice(Building container from scratch in Go)](https://www.youtube.com/watch?v=Utf-A4rODH8&list=WL&index=4&pp=gAQBiAQB)
