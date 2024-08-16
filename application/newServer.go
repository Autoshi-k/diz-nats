package application

import (
	"fmt"
)

const yamlPath = "./"

func (app DizNatsApp) NewServer(name, version, confPath string) (err error) {
	// new server --name --confp [path] or --conf [user/jwt] --override-port --override-ws-port

	_, ok := app.memory.LoadServers(name)
	if ok {
		return fmt.Errorf("server exist, choose new name")
	}

	path, err := app.memory.CreateNewFolder(name)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if nErr := app.memory.RemoveFolder(name); nErr != nil {
				err = fmt.Errorf("%w. failed to remove server %s", err, nErr.Error())
			}
		}
	}()

	// when we get to this function, version is already set (with latest or different version
	// 	confPath should already be set with default one or new one
	// todo should copy the conf to memory with server name for fast future use

	if err = app.memory.CopyFile(name, "server.conf", confPath); err != nil {
		return fmt.Errorf("failed to save conf file")
	}

	// todo write the docker-compose.yaml
	yaml, err := app.docker.GenerateDockerComposeYaml(name, path)

	if err = app.memory.SaveFile(name, "docker-compose.yaml", yaml); err != nil {
		return fmt.Errorf("failed to save conf file")
	}

	if err = app.docker.PullAndRun(version, path); err != nil {
		return err
	}

	// todo need to save the docker name and conf for future use and monitoring
	// 	also need to set log file

	// need to save all of this in memory somewhere
	// run "docker pull version" - if version empty then latest
	// run "docker run -p 4222:4222 -ti nats:latest"
	// or docker run --name nats --network nats --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222
	// or follow https://github.com/nats-io/nats.docs/issues/165 response
	return nil
}
