package utils

import (
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

var privateKeyFile = "/root/.ssh/id_rsa"

func getPublicKeys() (*ssh.PublicKeys, error) {
	if _, err := os.Stat(privateKeyFile); err != nil {
		log.Fatalf("Failed to find private key file: %s\n", err.Error())
		return nil, err
	}
	publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, "")
	if err != nil {
		log.Fatalf("generate publickeys failed: %s\n", err.Error())
		return nil, err
	}
	return publicKeys, nil
}

// CloneRepoToLocation will clone the repository to the specified path on
// the server
func CloneRepoToLocation(repository string, location string) error {
	publicKeys, err := getPublicKeys()
	if err != nil {
		return err
	}
	_, err = git.PlainClone(location, false, &git.CloneOptions{
		Auth:     publicKeys,
		URL:      repository,
		Progress: os.Stdout,
	})
	return err
}

// PullRepoAtLocation pulls latest changes from git repo
func PullRepoAtLocation(location string) error {
	repository, err := git.PlainOpen(location)
	if err != nil {
		log.Fatalf("Failed to open git path: %s\n", err.Error())
		return err
	}

	worktree, err := repository.Worktree()
	if err != nil {
		log.Fatalf("Failed to get worktree for repo: %s\n", err.Error())
		return err
	}

	publicKeys, err := getPublicKeys()
	if err != nil {
		return err
	}

	return worktree.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth:       publicKeys,
		Progress:   os.Stdout,
	})
}
