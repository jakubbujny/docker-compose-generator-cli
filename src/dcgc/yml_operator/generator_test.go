package yml_operator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStandardGeneration(t *testing.T) {
	assert := assert.New(t)
	//given
	ports := []string{"6666"}
	volumes := []string{"volumes"}
	image := "redis"

	//when
	output,_ := GenerateYml(ports,volumes,image,"")

	//then
	assert.Contains(output, "6666:6666", "Missing ports mapping in output")

}

func TestNamedVolumesGenerationSinglePath(t *testing.T) {
	assert := assert.New(t)
	//given
	path := []string{"/data"}
	service := "test"
	//when
	output := convertVolumesToNamedVolumes(path, service)
	//then
	assert.Contains(output[0], "test_data:/data")
}

func TestNamedVolumesGeneratioSameNameManyTimes(t *testing.T) {
	assert := assert.New(t)
	//given
	path := []string{"/data", "/data"}
	service := "test"
	//when
	output := convertVolumesToNamedVolumes(path, service)
	//then
	assert.Contains(output[0], "test_data:/data")
	assert.Contains(output[1], "test_data1:/data")
}

func TestNamedVolumesGenerationSinglePathEndingSlash(t *testing.T) {
	assert := assert.New(t)
	//given
	path := []string{"/data/"}
	service := "test"
	//when
	output := convertVolumesToNamedVolumes(path, service)
	//then
	assert.Contains(output[0], "test_data:/data")
}

func TestNamedVolumesGenerationMultipartPath(t *testing.T) {
	assert := assert.New(t)
	//given
	path := []string{"/data/some/dir"}
	service := "test"
	//when
	output := convertVolumesToNamedVolumes(path, service)
	//then
	assert.Contains(output[0], "test_dir:/data/some/dir")

}

func TestServiceNameFromImageWithoutNamespace(t *testing.T) {
	assert := assert.New(t)
	//given
	image := "jenkins"
	//when
	output := convertImageToServiceName(image)
	//then
	assert.Equal(output, "jenkins")
}

func TestServiceNameFromImageWithNamespace(t *testing.T) {
	assert := assert.New(t)
	//given
	image := "jenkinsNamespace/jenkinsName"
	//when
	output := convertImageToServiceName(image)
	//then
	assert.Equal(output, "jenkinsName")
}

func TestYmlMarshal(t *testing.T) {
	assert := assert.New(t)
	//given
	volumes := []string{"/data/"}
	ports := []string{"8080"}
	image := "jenkins"
	//when
	yml,error := GenerateYml(ports, volumes, image,"")
	//then
	assert.Nil(error)
	assert.Contains(yml, "8080:8080")
	assert.Contains(yml, "jenkins:")
	assert.Contains(yml, "data:/data")
}


func TestGenerateNamedVolumes(t *testing.T) {
	assert := assert.New(t)
	//given
	volumes := []string{"test_dir:/data/some/dir", "test_dir_some:/data/some/dir_some"}
	//when
	volumesString := generateNamedVolumesYml(volumes)
	//then
	assert.Equal("test_dir: {}\ntest_dir_some: {}\n", volumesString)
}

func TestInsertVolumeSectionWhenNotExists(t *testing.T) {
	assert := assert.New(t)

	//given
	source := "version:'3'\ntopLevel:"
	volumes := []string{"test_dir:/data/some/dir", "test_dir_some:/data/some/dir_some"}
	//when
	output := insertVolumesSection(source, volumes)
	//then

	assert.Equal("version:'3'\ntopLevel:\nvolumes:\n   test_dir: {}\n   test_dir_some: {}\n", output)
}

func TestInsertVolumeSectionWhenExists(t *testing.T) {
	assert := assert.New(t)

	//given
	source := "version:'3'\nvolumes:\ntopLevel:"
	volumes := []string{"test_dir:/data/some/dir", "test_dir_some:/data/some/dir_some"}
	//when
	output := insertVolumesSection(source, volumes)
	//then

	assert.Equal("version:'3'\nvolumes:\n   test_dir: {}\n   test_dir_some: {}\ntopLevel:\n", output)
}