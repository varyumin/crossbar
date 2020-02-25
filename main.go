package main

import (
	"fmt"
	"github.com/varyumin/crossbar/pkg/docker"
	"github.com/varyumin/crossbar/pkg/k8s"
)

func main() {
	var ns k8s.AllowNamespace
	userDocker := docker.DockerAuth{
		URL: "docker.fabric8.ru",
		Username: "varyumin",
		Password: "1130688W@r",
		Email: "fake@mail.net",
	}

	userDocker.


	images := ns.GetPodsImageList(ns.GetNsFromFile("namespaces.yaml"))
	fmt.Println(images)
}
