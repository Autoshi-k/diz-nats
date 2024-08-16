package memory

type Organizer interface {
	CopyFile(serverName, fileName, source string) error
	CreateNewFolder(serverName string) (path string, err error)
	LoadServers(names ...string) (servers interface{}, ok bool)
	RemoveFolder(serverName string) error
	SaveFile(serverName, fileName string, bytes []byte) error
}
