package docker

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type StatusDocker struct {
	Stream string `json:"stream"`
	Image  string `json:"image"`
}
type DockerImage struct {
	Domain    string
	ImageName string
	Tag       string
}

func (d *DockerImage) GetFullImageName() string {
	if d.Domain == "" {
		return fmt.Sprintf("%s", d.ImageName)
	} else {
		return fmt.Sprintf("%s/%s", d.Domain, d.ImageName)
	}

}
func (d *DockerImage) GetFullImageNameWithTag() string {
	if d.Domain == "" {
		return fmt.Sprintf("%s:%s", d.ImageName, d.Tag)
	} else {
		return fmt.Sprintf("%s/%s:%s", d.Domain, d.ImageName, d.Tag)
	}
}

func (d *DockerAuth) DockerPullWithAuth(image string) {
	ctx := context.Background()
	out, err := d.dockerClientConnect().ImagePull(
		ctx,
		image,
		types.ImagePullOptions{
			RegistryAuth: d.Login(),
		})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
}

func (d *DockerAuth) DockerPull(image string) {
	ctx := context.Background()
	out, err := d.dockerClientConnect().ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
}

func (d *DockerAuth) DockerPushWithAuth(image string) {
	ctx := context.Background()
	out, err := d.dockerClientConnect().ImagePush(ctx, image, types.ImagePushOptions{RegistryAuth: d.Login()})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
}
func (d *DockerAuth) DockerReTag(srcImage DockerImage, dstImage DockerImage) {
	ctx := context.Background()
	err := d.dockerClientConnect().ImageTag(ctx, srcImage.GetFullImageNameWithTag(), dstImage.GetFullImageNameWithTag())
	if err != nil {
		panic(err)
	}

}

func (d *DockerAuth) dockerClientConnect() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Panic(err)
	}
	return cli
}

func ParseImageTag(image string) DockerImage {
	imageStruct := DockerImage{}
	protoImageTags := strings.SplitN(image, ":", 2)
	if len(protoImageTags) == 1 {
		log.Errorf("unable to parse docker host `%s`", image)
	}
	imageName, tag := protoImageTags[0], protoImageTags[1]
	imageStruct.Tag = tag
	protoDomainImage := strings.SplitN(imageName, "/", 2)
	if len(protoDomainImage) == 1 {
		imageStruct.ImageName = imageName
	} else {
		imageStruct.Domain, imageStruct.ImageName = protoDomainImage[0], protoDomainImage[1]
	}
	return imageStruct
}

func (d *DockerAuth) DockerSaveToArchive(image string) {
	list_image := []string{}
	list_image = append(list_image, image)

	ctx := context.Background()
	out, err := d.dockerClientConnect().ImageSave(ctx, list_image)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	filePathStruc := ParseImageTag(image)
	filePath := fmt.Sprintf("%s_%s_%s",
		filePathStruc.Domain,
		strings.ReplaceAll(filePathStruc.ImageName, "/", "_"),
		filePathStruc.Tag)

	gzPath := fmt.Sprintf("%s.tar.gz", filePath)

	gz, err := os.Create(gzPath)

	if err != nil {
		log.Fatalln(err)
	}
	defer gz.Close()

	gw := gzip.NewWriter(gz)

	defer gw.Close()

	io.Copy(gw, out)
}

func (d *DockerAuth) DockerLoadFromArchive(path string) {
	imageArchive, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	out, err := d.dockerClientConnect().ImageLoad(ctx, imageArchive, true)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(out.Body)
	defer out.Body.Close()
	if err != nil {
		log.Panic(err)
	}
	var msg StatusDocker
	err = json.Unmarshal(b, &msg)
	fmt.Println(msg)

	fmt.Println(out.Body)
	io.Copy(os.Stdout, out.Body)
}
