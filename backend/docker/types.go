package docker

import (
	"fmt"

	"github.com/docker/go-connections/nat"
)

// ContainerMetadata represents all metadata required
// to be tracked for a container to be managed
type ContainerMetadata struct {
	ID   string
	Name string
	Port *nat.PortBinding
}

func (metadata ContainerMetadata) String() string {
	hostIP := "<none>"
	hostPort := "<none>"
	if metadata.Port != nil {
		hostIP = metadata.Port.HostIP
		hostPort = metadata.Port.HostPort
	}
	return fmt.Sprintf(`
	ContainerMetadata{
		ID: %s,
		Name: %s,
		Port: %s:%s
	}`, metadata.ID, metadata.Name, hostIP, hostPort)
}
