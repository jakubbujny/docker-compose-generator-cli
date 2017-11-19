package main

import (
	"os"
	"fmt"
	"dcgc/docker"
	"dcgc/yml_generator"
)

func main() {
	imageName := os.Args[2]
	docker := docker.New()
	docker.RemoveImage(imageName)
	docker.PullImage(imageName)
	ports,volumes,_ := docker.InspectImage(imageName)
	ymlFile, _ := yml_generator.GenerateYml(ports,volumes,imageName)
	fmt.Println(ymlFile)
}
