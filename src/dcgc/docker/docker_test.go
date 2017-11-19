package docker

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRemoveImage(t *testing.T) {
	assert := assert.New(t)
	//given
	docker := New()
	image := "library/busybox"
	//when
	res,err := docker.RemoveImage(image)
	assert.Nil(err)
	assert.Contains(res, "busybox:latest")
}

func TestRemoveImageNotFound(t *testing.T) {
	assert := assert.New(t)
	//given
	docker := New()
	image := "library/busyboxunknown"
	//when
	_,err := docker.RemoveImage(image)
	//then
	assert.NotNil(err)
	assert.Contains(err.Error(), "No such image")

}

func TestPullImageNotFoundWhilePulling(t *testing.T) {
	assert := assert.New(t)
	//given
	docker := New()
	image := "busyboxunknown"
	//when
	_,err := docker.PullImage(image)
	//then
	assert.NotNil(err)
	assert.Contains(err.Error(), "pull access denied for busyboxunknown")

}

func TestPullNormalImageWihtoutNamespace(t *testing.T) {
	assert := assert.New(t)
	//given
	docker := New()
	image := "library/busybox"
	//when
	res,err := docker.PullImage(image)
	assert.Nil(err)
	assert.Contains(res, "Pulling from library/busybox")

}

func TestPullNormalImageWithNamespace(t *testing.T) {
	assert := assert.New(t)
	//given
	docker := New()
	image := "gliderlabs/alpine"
	//when
	res,err := docker.PullImage(image)
	assert.Nil(err)
	assert.Contains(res, "Pulling from gliderlabs/alpine")

}

func TestInspectImageWhenEverythingIs(t *testing.T) {
	assert := assert.New(t)
	//given
	docker := New()
	image := "redis"
	//when
	docker.PullImage(image)
	ports,volumes,err := docker.InspectImage(image)
	assert.Nil(err)
	assert.True(len(ports) == 1)
	assert.True(len(volumes) == 1)

}

func TestInspectImageWhenMissing(t *testing.T) {
	assert := assert.New(t)
	//given
	docker := New()
	image := "python"
	//when
	docker.PullImage(image)
	ports,volumes,err := docker.InspectImage(image)
	assert.Nil(err)
	assert.True(len(ports) == 0)
	assert.True(len(volumes) == 0)

}