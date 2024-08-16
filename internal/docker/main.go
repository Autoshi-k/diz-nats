package docker

type Docker interface {
	GenerateDockerComposeYaml(name, path string) (file []byte, err error)
	PullAndRun(version, path string) error
}
