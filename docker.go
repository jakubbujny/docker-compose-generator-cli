package main

import (
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"context"
	"bytes"
)
type Docker struct {
	cli *client.Client
}
func makeDocker() Docker {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return Docker{cli:cli}
}

func (docker *Docker) pullImage(imageName string) (string, error) {
	result,err := docker.cli.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		return "", err
	} else {
		defer result.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(result)
		return buf.String(), nil
	}

}

func (docker *Docker) removeImage(imageName string) (string, error) {
	result,err := docker.cli.ImageRemove(context.Background(), imageName, types.ImageRemoveOptions{})
	if err != nil {
		return "", err
	} else {
		toReturn  := ""
		for _,element := range result {
			toReturn += element.Untagged +";"
		}
		return toReturn, err
	}

}