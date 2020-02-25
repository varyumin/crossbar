package docker

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
)

type DockerAuth struct {
	URL      string
	Username string
	Password string
	Email    string
}

func (d *DockerAuth) Login() string {
	authConfig := types.AuthConfig{
		Username: d.Username,
		Password: d.Password,
		Email:    d.Email,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.URLEncoding.EncodeToString(encodedJSON))
	return base64.URLEncoding.EncodeToString(encodedJSON)
}
