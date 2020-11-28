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
	return fmt.Sprintf(`
	ContainerMetadata{
		ID: %s,
		Name: %s,
		Port: %s:%s
	}`, metadata.ID, metadata.Name, metadata.Port.HostIP, metadata.Port.HostPort)
}
