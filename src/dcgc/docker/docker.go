package docker

import (
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"context"
	"bytes"
)
type Docker struct {
	cli *client.Client
}
func New() Docker {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return Docker{cli:cli}
}

func (docker *Docker) PullImage(imageName string) (string, error) {
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

func (docker *Docker) InspectImage(imageName string) ([]string,[]string,error) {
	imageInspect,_,err := docker.cli.ImageInspectWithRaw(context.Background(), imageName)
	if err != nil {
		return nil,nil, err
	} else {
		ports := make([]string,0, len(imageInspect.Config.ExposedPorts))
		for k,_ := range imageInspect.Config.ExposedPorts {
			ports = append(ports,k.Port())
		}
		volumes := make([]string,0, len(imageInspect.Config.Volumes))
		for k,_ := range imageInspect.Config.Volumes {
			volumes = append(volumes,k)
		}
		return ports,volumes,nil
	}
}

func (docker *Docker) RemoveImage(imageName string) (string, error) {
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