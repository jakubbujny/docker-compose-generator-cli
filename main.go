package main

import (
	"os"
	"fmt"
)

func main() {
	imageName := os.Args[2]
	docker := MakeDocker()
	docker.RemoveImage(imageName)
	docker.PullImage(imageName)
	ports,volumes,_ := docker.InspectImage(imageName)
	ymlFile, _ := GenerateToolService(ports,volumes,imageName)
	fmt.Println(ymlFile)
}
