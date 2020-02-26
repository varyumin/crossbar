package main

import (
	"github.com/varyumin/crossbar/pkg/docker"
)

func main() {
	//var ns k8s.AllowNamespace
	userDocker := docker.DockerAuth{
		URL:      "docker.fabric8.ru",
		Username: "varyumin",
		Password: "1130688W@r",
		Email:    "fake@mail.net",
	}

	userDocker.Login()

	//images := ns.GetPodsImageList(ns.GetNsFromFile("namespaces.yaml"))
	//for _, img :=range images{
	//	fmt.Println(img)
	//}
	//fmt.Println(images)
}
