package yml_operator

import (
	"gopkg.in/yaml.v2"
	"strings"
	"regexp"
	"strconv"
)

type DockerService struct {
	Image   string   `yaml:"image"`
	Ports   []string `yaml:"ports"`
	Volumes []string `yaml:"volumes"`
}
func convertPortsToPublishForm(ports []string) []string {
	toReturn := make([]string, 0, len(ports))
	for _,val := range ports {
		toReturn = append(toReturn, val + ":" + val)
	}
	return toReturn
}
func existsInList(list []string, element string) bool {
	for _,el := range list {
		if strings.Contains(el, element) {
			return true
		}
	}
	return false
}
func convertVolumesToNamedVolumes(volumes []string, serviceName string) []string {
	toReturn := make([]string, 0, len(volumes))

	for _, volumePath := range volumes {
		sufix := ""
		if len(regexp.MustCompile("/").FindAllStringIndex(volumePath, -1)) > 1 && !strings.HasSuffix(volumePath,"/") {
			split := strings.Split(volumePath, "/")
			sufix = split[len(split)-1]
		} else {
			sufix = strings.Replace(volumePath, "/", "", -1)
		}

		volumeName := serviceName + "_" + sufix
		i := 1
		for existsInList(toReturn, volumeName) {
			volumeName = serviceName+"_" + sufix + strconv.Itoa(i)
			i++
		}
		dockerVolumeString := volumeName + ":" + volumePath

		toReturn = append(toReturn, dockerVolumeString)
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

func generateNamedVolumesYml(volumes []string) string {
	output := ""
	for _,volume := range volumes {
		output += strings.Split(volume, ":")[0]+": {}" + "\n"
	}
	return output
}
func insertVolumesSection(sourceYml string, volumes []string) string {
	volumesString := generateNamedVolumesYml(volumes)
	if !strings.Contains(sourceYml, "volumes:") {
		sourceYml += "\nvolumes: placeholder"
	}
	output,_ := AppendToYmlInSection(volumesString, sourceYml, "volumes")
	return output
}
func GenerateYml(ports []string, volumes []string, image string, sourceYml string) (string, error) {
	if sourceYml == "" {
		sourceYml = "version: '3'\nservices:"
	}
	serviceName := convertImageToServiceName(image)
	sourceYml = insertVolumesSection(sourceYml, convertVolumesToNamedVolumes(volumes, serviceName))
	service := DockerService{image, convertPortsToPublishForm(ports), convertVolumesToNamedVolumes(volumes, serviceName)}

	serviceYml, err := yaml.Marshal(map[string]interface{} {
		serviceName: service,
	})
	if err != nil {
		return "", err
	}
	outputYml, err := AppendToYmlInSection(string(serviceYml),sourceYml, "services")
	if err != nil {
		return "", err
	}
	return outputYml, nil
}
