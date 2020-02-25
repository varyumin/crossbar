package main

import (
	"fmt"
	"github.com/varyumin/crossbar/pkg/k8s"
)

func main() {
	var ns k8s.AllowNamespace
	images := ns.GetPodsImageList(ns.GetNsFromFile("namespaces.yaml"))
	for _, img := range images {
		fmt.Println(img)
	}
}
