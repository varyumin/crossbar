package main

import (
	"fmt"
	"github.com/varyumin/crossbar/pkg/docker"
	"github.com/varyumin/crossbar/pkg/k8s"
)

func Pull(images []string, dockerLogIn *docker.DockerAuth) {
	for _, image := range images {
		//log.Infof("Docker pull image: %s", image)
		sepImage := docker.ParseImageTag(image)
		fmt.Printf("Full: %s Domen: %s  Image: %s Tag: %s \n", image, sepImage.Domain, sepImage.ImageName, sepImage.Tag)
		//dockerLogIn.DockerPullWithAuth(image)
	}
}

func main() {
	var ns k8s.AllowNamespace
	userDocker := docker.DockerAuth{
		URL:      "docker.fabric8.ru",
		Username: "vryumin",
		Password: "1130688W@r",
		Email:    "fake@mail.net",
	}
	userDocker.Login()

	images := ns.GetPodsImageList(ns.GetNsFromFile("namespaces.yaml"))
	Pull(images, &userDocker)
}
