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
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/mitchellh/go-homedir"
)

var dockerClient *client.Client = nil
var ctx context.Context = nil

func getBuildContext(buildPath string) (io.Reader, error) {
	ctx, err := archive.TarWithOptions(buildPath, &archive.TarOptions{})
	if err != nil {
		log.Printf("Failed to tar build directory: %s\n", err.Error())
		return nil, err
	}
	return ctx, err
}

// Connect initializes docker client
func Connect() error {
	var err error
	ctx = context.Background()
	dockerClient, err = client.NewEnvClient()
	return err
}

// BuildImage builds docker image from dockerfile path passed in
func BuildImage(dockerfile string, imageName string, imageTag string) error {
	dockerfile, _ = homedir.Expand(dockerfile)
	dir := path.Dir(dockerfile)
	dockerfile = path.Base(dockerfile)
	buildContext, err := getBuildContext(dir)
	if err != nil {
		return err
	}

	cleanedName := cleanImageName(imageName)
	resp, err := dockerClient.ImageBuild(ctx, buildContext, types.ImageBuildOptions{
		Tags:           []string{cleanedName + ":" + imageTag},
		Dockerfile:     dockerfile,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
		SuppressOutput: false,
	})
	if err != nil {
		log.Printf("Failed to build docker image: %s\n", err.Error())
		return err
	}

	return readStreamForStatus(resp.Body)
}

// StartContainer starts a new container with the specified image
func StartContainer(imageName string, imageTag string, uniqueID string) (*ContainerMetadata, error) {
	cleanedName := cleanImageName(imageName)
	containerConfig := container.Config{
		Image: cleanedName + ":" + imageTag,
	}
	hostConfig := container.HostConfig{
		PublishAllPorts: true,
	}

	containerMetadata := ContainerMetadata{
		Name: cleanedName + uniqueID,
	}
	created, err := dockerClient.ContainerCreate(ctx, &containerConfig, &hostConfig, nil, containerMetadata.Name)
	if err != nil {
		return nil, err
	}
	containerMetadata.ID = created.ID

	err = dockerClient.ContainerStart(ctx, containerMetadata.ID, types.ContainerStartOptions{})
	if err != nil {
		return nil, err
	}

	inspection, err := dockerClient.ContainerInspect(ctx, containerMetadata.ID)
	if err != nil {
		return nil, err
	}

	for port := range inspection.NetworkSettings.Ports {
		containerMetadata.Port = &inspection.NetworkSettings.Ports[port][0]
		break
	}

	return &containerMetadata, nil
}

// DeleteContainer force removes container with given ID
func DeleteContainer(containerID string) error {
	return dockerClient.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
		Force: true,
	})
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

func cleanImageName(imageName string) string {
	return strings.ToLower(strings.TrimSpace(imageName))
}
