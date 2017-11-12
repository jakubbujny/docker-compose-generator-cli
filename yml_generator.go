package main

import (
	"github.com/ghodss/yaml"
	"strings"
	"regexp"
)

type DockerService struct {
	Image   string   `json:"image"`
	Ports   []string `json:"ports"`
	Volumes []string `json:"volumes"`
}
func convertPortsToPublishForm(ports []string) []string {
	toReturn := make([]string, 0, len(ports))
	for _,val := range ports {
		toReturn = append(toReturn, val + ":" + val)
	}
	return toReturn
}
func convertVolumesToNamedVolumes(volumes []string, serviceName string) []string {
	toReturn := make([]string, 0, len(volumes))

	for _,volume := range volumes {
		sufix := ""
		if len(regexp.MustCompile("/").FindAllStringIndex(volume, -1)) > 1 && !strings.HasSuffix(volume,"/") {
			split := strings.Split(volume, "/")
			sufix = split[len(split)-1]
		} else {
			sufix = strings.Replace(volume, "/", "", -1)
		}
		toReturn = append(toReturn, serviceName+"_"+sufix+":"+volume)
	}
	return toReturn
}
func convertImageToServiceName(imageName string) string {
	if strings.Contains(imageName, "/") {
		return strings.Split(imageName, "/")[1]
	} else {
		return imageName
	}
}
func GenerateToolService(ports []string, volumes []string, image string) (string, error) {
	serviceName := convertImageToServiceName(image)
	service := DockerService{image, convertPortsToPublishForm(ports), convertVolumesToNamedVolumes(volumes, serviceName)}

	o, err := yaml.Marshal(map[string]interface{} {
		serviceName: service,
	})
	if err != nil {
		return "", err
	}
	return string(o), nil
}
