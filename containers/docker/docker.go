package docker

import (
	// "os"

	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

var DockerClient *client.Client

func StartDockerClient() error {
	var err error
	DockerClient, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Info("Fail Start Docker Client")
		log.Fatal(err)
		return err
	}
	log.Info("Started Docker Client")
	return nil

}