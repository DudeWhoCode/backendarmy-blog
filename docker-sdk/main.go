package main

import (
	"github.com/dudewhocode/backendarmy-blog/docker-sdk/container"
)

func main() {
	container.CreateNewContainer("nginx:latest")
}
