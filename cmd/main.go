package main

import (
	"diz-nats/application"
	"diz-nats/infrastructure/docker"
	"diz-nats/infrastructure/memory"
)

func main() {
	application.NewApp(memory.NewMemoryOrganizer(), docker.NewDocker())
}

// todo next commands:
// list servers/listeners/mockers/confs/name
// 		flags --open --closed
// 		list without any following shit will show servers/listeners/mockers and status
// new server/listener/mocker/conf
// 		new server --name --confp [path] or --conf [user/jwt] --override-port --override-ws-port
// 		new listener --name [name] --server [server-name] --channel --channels
// 		new mocker --name [name] --server [server-name] --channel [subject] --cnd [checks if key in msg.Data] --resp [interface]
//		new conf
