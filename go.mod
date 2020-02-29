module github.com/varyumin/crossbar

go 1.12

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Azure/go-autorest/autorest v0.10.0 // indirect
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/containerd/containerd v1.3.3 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v0.0.0-20200226101133-26e6708ccf63
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/elazarl/goproxy v0.0.0-20180725130230-947c36da3153 // indirect
	github.com/gophercloud/gophercloud v0.8.0 // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/onsi/ginkgo v1.11.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/prometheus/common v0.9.1
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/net v0.0.0-20200222125558-5a598a2470a0 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	gopkg.in/yaml.v2 v2.2.8
	gotest.tools/v3 v3.0.2 // indirect
	k8s.io/api v0.17.3 // indirect
	k8s.io/apimachinery v0.17.3
	k8s.io/client-go v0.0.0-20200222082411-f772958f8a29
	k8s.io/kube-openapi v0.0.0-20200121204235-bf4fb3bd569c // indirect
	k8s.io/utils v0.0.0-20200124190032-861946025e34 // indirect
	sigs.k8s.io/structured-merge-diff/v3 v3.0.0-20200207200219-5e70324e7c1c // indirect
)

replace github.com/varyumin/crossbar/pkg/docker v0.0.0 => ./pkg/docker

replace github.com/varyumin/crossbar/pkg/k8s v0.0.0 => ./pkg/k8s
