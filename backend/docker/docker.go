package docker

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"path"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/mitchellh/go-homedir"
)

func getBuildContext(buildPath string) (io.Reader, error) {
	ctx, err := archive.TarWithOptions(buildPath, &archive.TarOptions{})
	if err != nil {
		log.Fatalf("Failed to tar build directory: %s\n", err.Error())
		return nil, err
	}
	return ctx, err
}

// BuildImage builds docker image from dockerfile path passed in
func BuildImage(dockerfile string, imageName string) error {
	ctx := context.Background()
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Unable to init client: %s\n", err.Error())
		return err
	}
	dockerfile, _ = homedir.Expand(dockerfile)
	dir := path.Dir(dockerfile)
	dockerfile = path.Base(dockerfile)
	buildContext, err := getBuildContext(dir)
	if err != nil {
		return err
	}

	cleanedName := strings.ToLower(strings.TrimSpace(imageName))
	resp, err := dockerClient.ImageBuild(ctx, buildContext, types.ImageBuildOptions{
		Tags:           []string{cleanedName},
		Dockerfile:     dockerfile,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
		SuppressOutput: false,
	})
	if err != nil {
		log.Fatalf("Failed to build docker image: %s\n", err.Error())
		return err
	}

	return readStreamForStatus(resp.Body)
}

func readStreamForStatus(daemonStream io.ReadCloser) error {
	defer daemonStream.Close()

	decoder := json.NewDecoder(daemonStream)
	for {
		var jsonMessage jsonmessage.JSONMessage
		if err := decoder.Decode(&jsonMessage); err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error reading docker daemon messages: %s\n", err.Error())
			return err
		}
		if jsonMessage.Error != nil {
			return errors.New(jsonMessage.ErrorMessage)
		}
	}
	return nil
}
