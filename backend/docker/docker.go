package docker

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
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
func BuildImage(dockerfile string, imageName string) (types.ImageBuildResponse, error) {
	ctx := context.Background()
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Unable to init client: %s\n", err.Error())
		return types.ImageBuildResponse{}, err
	}
	dockerfile, _ = homedir.Expand(dockerfile)
	dir := path.Dir(dockerfile)
	dockerfile = path.Base(dockerfile)
	buildContext, err := getBuildContext(dir)
	if err != nil {
		return types.ImageBuildResponse{}, err
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
		return resp, err
	}

	termFd, isTerm := term.GetFdInfo(os.Stderr)
	jsonmessage.DisplayJSONMessagesStream(resp.Body, os.Stderr, termFd, isTerm, tagImage)

	return resp, nil
}

func tagImage(aux *json.RawMessage) {
	var result types.BuildResult
	if err := json.Unmarshal(*aux, &result); err != nil {
		log.Fatalf("Failed to unmarshal aux message. Cause: %s", err)
	}
	log.Printf("result.ID: %s\n", result.ID)
}

func parseDockerDaemonJSONMessages(daemonMessage io.Reader) (types.BuildResult, error) {
	decoder := json.NewDecoder(daemonMessage)
	var jsonMessage jsonmessage.JSONMessage
	if err := decoder.Decode(&jsonMessage); err != nil {
		return types.BuildResult{}, err
	}
	if err := jsonMessage.Error; err != nil || jsonMessage.Aux == nil {
		log.Print("Incorrectly formatted json")
		log.Print(jsonMessage.Status)
		return types.BuildResult{}, err
	}
	var result types.BuildResult
	if err := json.Unmarshal(*jsonMessage.Aux, &result); err != nil {
		log.Fatalf("Failed to unmarshal aux message. Cause: %s", err)
		return types.BuildResult{}, err
	}
	return result, nil
}
