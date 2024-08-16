package nats

import (
	"fmt"
	"log"
	"os/exec"
)

const yamlPath = "./"

// todo just like that ?....
// new server --name --confp [path] or --conf [user/jwt] --override-port --override-ws-port
func NewServer(name, version, confPath string) {
	// todo validate name in memory
	// 	am i already using this name?

	// when we get to this function, version is already set (with latest or different version
	// 	confPath should already be set with default one or new one
	// todo should copy the conf to memory with server name for fast future use

	// wrap...
	cmd := exec.Command("docker pull " + version) // version after validation and modification if needed
	if err := cmd.Run(); err != nil {
		log.Fatal(err) // should return the error
	}

	// todo write the docker-compose.yaml

	// run docker-compose -f docker-compose.yaml up
	cmd = exec.Command(fmt.Sprintf("run docker-compose -f %s/docker-compose.yaml up", yamlPath))
	if err := cmd.Run(); err != nil {
		log.Fatal(err) // should return the error
	}

	// todo need to save the docker name and conf for future use and monitoring
	// 	also need to set log file

	// need to save all of this in memory somewhere
	// run "docker pull version" - if version empty then latest
	// run "docker run -p 4222:4222 -ti nats:latest"
	// or docker run --name nats --network nats --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222
	// or follow https://github.com/nats-io/nats.docs/issues/165 response

}
