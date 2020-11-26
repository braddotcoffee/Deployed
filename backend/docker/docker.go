package docker

import (
	"context"
	"io"
	"log"
	"path"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
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
func BuildImage(dockerfile string) (types.ImageBuildResponse, error) {
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

	resp, err := dockerClient.ImageBuild(ctx, buildContext, types.ImageBuildOptions{
		Dockerfile: dockerfile,
	})
	if err != nil {
		log.Fatalf("Failed to build docker image: %s\n", err.Error())
		return resp, err
	}
	return resp, nil
}
