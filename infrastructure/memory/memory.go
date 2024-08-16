package memory

import (
	"diz-nats/internal/memory"
	"fmt"
	"io/ioutil"
	"os"
)

type MemOrganizer struct {
}

// todo move to struct and allow changing the root location
const root = "./temp/"

func (m MemOrganizer) CopyFile(serverName, fileName, source string) error {
	// Read all content of src to data, may cause OOM for a large file.
	data, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}
	// Write data to dst
	return ioutil.WriteFile(root+serverName+"/"+fileName, data, 0644)
}

func (m MemOrganizer) CreateNewFolder(serverName string) (path string, err error) {
	if err = os.Mkdir(root+serverName, os.ModePerm); err != nil {
		return "", err
	} else {
		return root + serverName, nil
	}
}

func (m MemOrganizer) LoadServers(names ...string) (servers interface{}, ok bool) {
	d, err := os.Open(root)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	defer d.Close()

	folders, err := d.Readdirnames(-1)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	var found int
	for name := range names {
		for folderName := range folders {
			if name == folderName {
				found++
			}
		}
	}

	// why do I even return anything?
	return nil, found == len(names)
}

func (m MemOrganizer) RemoveFolder(serverName string) error {
	return os.RemoveAll(root + serverName)
}

func (m MemOrganizer) SaveFile(serverName, fileName string, bytes []byte) error {
	return ioutil.WriteFile(root+serverName+"/"+fileName, bytes, 0644)
}

func NewMemoryOrganizer() memory.Organizer {
	return MemOrganizer{}
}
