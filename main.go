package main

import (
	"dcgc/docker"
	"dcgc/yml_operator"
	"github.com/urfave/cli"
	"os"
	"io/ioutil"
)

var DockerComposeFileName = "docker-compose.yml"

func addTool(toolName string) {
	docker := docker.New()
	docker.RemoveImage(toolName)
	_,error := docker.PullImage(toolName)
	if error != nil {
		panic("Problem while pulling docker image: "+error.Error())
	}
	ports,volumes,_ := docker.InspectImage(toolName)
	sourceYml := ""
	if _, err := os.Stat(DockerComposeFileName); !os.IsNotExist(err) {
		dockerComposeConfigRaw, err := ioutil.ReadFile(DockerComposeFileName)
		if err != nil {
			panic("There was problem while reading you docker-compose.yml file: " +err.Error())
		} else {
			sourceYml = string(dockerComposeConfigRaw)
		}
	}
	ymlFile, _ := yml_operator.GenerateYml(ports,volumes,toolName, sourceYml)

	err := ioutil.WriteFile(DockerComposeFileName, []byte(ymlFile), 0644)
	if err != nil {
		panic("There was problem while saving docker-compose.yml file: " + err.Error())
	}
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
