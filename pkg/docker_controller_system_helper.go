package pkg

import (
	"context"
	"github.com/docker/docker/client"
)

func getDockerRootDir(client *client.Client) (string, error) {
	info, err := client.Info(context.Background())
	if err != nil {
		return "", err
	}

	return info.DockerRootDir, nil
}