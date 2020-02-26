package main

import (
	"github.com/prometheus/common/log"
	"github.com/varyumin/crossbar/pkg/docker"
	"github.com/varyumin/crossbar/pkg/k8s"
)

func Pull(images []string, dockerLogIn *docker.DockerAuth) {
	for _, image := range images {
		log.Infof("Docker pull image: %s", image)
		dockerLogIn.DockerPull(image)
	}
}

func main() {
	var ns k8s.AllowNamespace
	userDocker := docker.DockerAuth{
		URL:      "docker.fabric8.ru",
		Username: "varyumin",
		Password: "1130688W@r",
		Email:    "fake@mail.net",
	}
	userDocker.Login()

	images := ns.GetPodsImageList(ns.GetNsFromFile("namespaces.yaml"))
	Pull(images, &userDocker)
}
