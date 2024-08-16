package docker

import (
	"diz-nats/internal/docker"
	"fmt"
	"os/exec"
)

type Docker struct {
}

func (d Docker) GenerateDockerComposeYaml(name, path string) (file []byte, err error) {
	return []byte(`
version: "1"
services:
  nats:
    image: nats
    command: "-c ` + path + `/server.conf
    volumes:
       - ./config/:/etc/nats
networks:
  default:
    external:
      name:` + name), nil
}

func (d Docker) PullAndRun(version, path string) error {
	cmd := exec.Command("docker pull nats" + version) // version after validation and modification if needed
	if err := cmd.Run(); err != nil {
		return err
	}

	// run docker-compose -f docker-compose.yaml up
	cmd = exec.Command(fmt.Sprintf("run docker-compose -f %s/docker-compose.yaml up", path))
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func NewDocker() docker.Docker {
	return Docker{}
}
