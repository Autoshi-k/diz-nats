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
	cmdStr := fmt.Sprintf("docker pull nats %s", version)
	_, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	if err != nil {
		return err
	}

	// run docker-compose -f docker-compose.yaml up
	cmdStr = fmt.Sprintf("run docker-compose -f %s/docker-compose.yaml up", path)
	_, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
	if err != nil {
		return err
	}

	return nil
}

func NewDocker() docker.Docker {
	return Docker{}
}
