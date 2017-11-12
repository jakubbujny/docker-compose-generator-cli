package main

import (
	"testing"
	"strings"
)

func TestRemoveImage(t *testing.T) {
	docker := makeDocker()
	res,err := docker.removeImage("library/busybox")
	if err != nil  {
		t.Fatal("It's ok so it should be nil")
	}
	if !strings.Contains(res,"busybox:latest") {
		t.Fatal("It should be deleted in log but have: "+res)
	}
}

func TestPullImageNotFoundWhilePulling(t *testing.T) {
	docker := makeDocker()
	_,err := docker.pullImage("library/busyboxnot")
	if err == nil  {
		t.Fatal("It's error so it cannot be nil")
	}
	if ! strings.Contains(err.Error(), "pull access denied for busyboxnot") {
		t.Error("Should be access denied in output but is: "+err.Error())
	}
}

func TestPullNormalImageWihtoutNamespace(t *testing.T) {
	docker := makeDocker()

	res,err := docker.pullImage("library/busybox")
	if err != nil  {
		t.Fatal("It's ok so it should be nil")
	}
	if !strings.Contains(res, "Pulling from library/busybox") {
		t.Error("Not found pulling in response but have: "+res)
	}
}