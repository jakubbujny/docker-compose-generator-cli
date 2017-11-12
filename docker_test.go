package main

import (
	"testing"
	"strings"
)

func TestRemoveImage(t *testing.T) {
	docker := MakeDocker()
	res,err := docker.RemoveImage("library/busybox")
	if err != nil  {
		t.Fatal("It's ok so it should be nil")
	}
	if !strings.Contains(res,"busybox:latest") {
		t.Fatal("It should be deleted in log but have: "+res)
	}
}

func TestRemoveImageNotFound(t *testing.T) {
	docker := MakeDocker()
	_,err := docker.RemoveImage("library/busyboxnot")
	if err == nil  {
		t.Fatal("It's invalid so it shouldn't be nil")
	}
	if ! strings.Contains(err.Error(), "No such image") {
		t.Error("Should be No such IMAGE but is: "+err.Error())
	}
}

func TestPullImageNotFoundWhilePulling(t *testing.T) {
	docker := MakeDocker()
	_,err := docker.PullImage("library/busyboxnot")
	if err == nil  {
		t.Fatal("It's error so it cannot be nil")
	}
	if ! strings.Contains(err.Error(), "pull access denied for busyboxnot") {
		t.Error("Should be access denied in output but is: "+err.Error())
	}
}

func TestPullNormalImageWihtoutNamespace(t *testing.T) {
	docker := MakeDocker()
	res,err := docker.PullImage("library/busybox")
	if err != nil  {
		t.Fatal("It's ok so it should be nil")
	}
	if !strings.Contains(res, "Pulling from library/busybox") {
		t.Error("Not found pulling in response but have: "+res)
	}
}

func TestPullNormalImageWithNamespace(t *testing.T) {
	docker := MakeDocker()

	res,err := docker.PullImage("gliderlabs/alpine")
	if err != nil  {
		t.Fatal("It's ok so it should be nil")
	}
	if !strings.Contains(res, "Pulling from gliderlabs/alpine") {
		t.Error("Not found pulling in response but have: "+res)
	}
}

func TestInspectImageWhenEverythingIs(t *testing.T) {
	docker := MakeDocker()

	docker.PullImage("redis")
	ports,volumes,err := docker.InspectImage("redis")
	if err != nil  {
		t.Fatal("It's ok so it should be nil")
	}
	if len(ports) != 1 {
		t.Error("Probably something is wrong when we see there more than one port")
	}
	if len(volumes) != 1 {
		t.Error("Probably something is wrong when we see there more than one volume")
	}

}

func TestInspectImageWhenMissing(t *testing.T) {
	docker := MakeDocker()

	docker.PullImage("python")
	ports,volumes,err := docker.InspectImage("python")
	if err != nil  {
		t.Fatal("It's ok so it should be nil")
	}
	if len(ports) != 0 {
		t.Error("Probably something is wrong when we see there more than zero port")
	}
	if len(volumes) != 0 {
		t.Error("Probably something is wrong when we see there more than zero volume")
	}

}