package main

import (
	"fmt"
	"dcgc/docker"
	"dcgc/yml_operator"
	"github.com/urfave/cli"
	"os"
)

var DockerComposeFileName = "docker-compose.yml"

func addTool(toolName string) {
	docker := docker.New()
	docker.RemoveImage(toolName)
	docker.PullImage(toolName)
	ports,volumes,_ := docker.InspectImage(toolName)
	ymlFile, _ := yml_operator.GenerateYml(ports,volumes,toolName)
	if _, err := os.Stat(DockerComposeFileName); os.IsExist(err) {

	} else {

	}
	fmt.Println(ymlFile)
}

func main() {
	app := cli.NewApp()
	app.Name = "docker-compose-generator-cli"
	app.Usage = "tool for fast prototyping using docker images from dockerhub generating docker-compose configs"
	app.Commands = []cli.Command{
		{
			Name:    "tool",
			Aliases: []string{"t"},
			Usage:   "add tool from dockerhub into docker-compose config",
			Action: func(c *cli.Context) error {
				addTool(c.Args().First())
				return nil
			},
		},
	}
	app.Run(os.Args)
}
