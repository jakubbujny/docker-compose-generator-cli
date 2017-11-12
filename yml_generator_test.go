package main

import (
	"testing"
	"strings"
)

func TestStandardGeneration(t *testing.T) {
	output,_ := GenerateToolService([]string{"6666"},[]string{"volumes"}, "redis" )
	if !strings.Contains(output, "6666:6666") {
		t.Error("Missing ports mapping in output")
	}
}

func TestNamedVolumesGenerationSinglePath(t *testing.T) {
	path := "/data"
	output := convertVolumesToNamedVolumes([]string{path}, "test")
	if !strings.Contains(output[0], "test_data:/data") {
		t.Error("Failed generation, output: "+output[0])
	}
}

func TestNamedVolumesGenerationSinglePathEndingSlash(t *testing.T) {
	path := "/data/"
	output := convertVolumesToNamedVolumes([]string{path}, "test")
	if !strings.Contains(output[0], "test_data:/data") {
		t.Error("Failed generation, output: "+output[0])
	}
}

func TestNamedVolumesGenerationMultipartPath(t *testing.T) {
	path := "/data/some/dir"
	output := convertVolumesToNamedVolumes([]string{path}, "test")
	if !strings.Contains(output[0], "test_data:/data/some/dir") {
		t.Error("Failed generation, output: "+output[0])
	}
}

func TestServiceNameFromImageWithoutNamespace(t *testing.T) {
	image := "jenkins"
	output := convertImageToServiceName(image)
	if output != "jenkins" {
		t.Error("Failed generation, output: "+output)
	}
}

func TestServiceNameFromImageWithNamespace(t *testing.T) {
	image := "jenkinsNamespace/jenkinsName"
	output := convertImageToServiceName(image)
	if output != "jenkinsName" {
		t.Error("Failed generation, output: "+output)
	}
}