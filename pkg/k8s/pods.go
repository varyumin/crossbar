package k8s

import (
	"context"
	"github.com/prometheus/common/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
)

type AllowNamespace struct {
	Namespaces []struct {
		Namespace string `yaml:"namespace"`
	} `yaml:"namespaces"`
}

func (a *AllowNamespace) GetNsFromFile(path string) AllowNamespace {
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		log.Infof("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, a)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return *a
}


func (a *AllowNamespace) GetPodsImageList(ns AllowNamespace)  []string{
	var images []string
	k8s := GetConnect()
	for _,v := range ns.Namespaces{
		pods, err := k8s.CoreV1().Pods(v.Namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		for _, v := range pods.Items {
			for _, p := range v.Spec.Containers{
				images = append(images, p.Image)
			}
		}
	}
	sort.Strings(images)
	return unique(images)
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}