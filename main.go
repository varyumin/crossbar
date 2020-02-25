package main

import (
	"fmt"
	"github.com/varyumin/crossbar/pkg/docker"
	"github.com/varyumin/crossbar/pkg/k8s"
)

func main() {
	var ns k8s.AllowNamespace
	var userDocker docker.DockerAuth

	images := ns.GetPodsImageList(ns.GetNsFromFile("namespaces.yaml"))
	fmt.Println(images)
}
