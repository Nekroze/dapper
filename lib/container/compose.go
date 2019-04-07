package container

import (
	"fmt"

	"github.com/docker/libcompose/config"
	"github.com/docker/libcompose/project"
)

type ComposeBackend struct {
	project.Project
}

func InitCompose(app Application) Backend {
	return &ComposeBackend{}.init(app)
}

func (b *ComposeBackend) init() *ComposeBackend {
	b.Name = "dapper"
	b.ServiceConfigs = config.NewServiceConfigs()
	for name, tainer := range app {
		b.ServiceConfigs.Add(
			name,
			b.containerToService(tainer),
		)
	}
	return b
}

func (b *ComposeBackend) containerToService(tainer Container) *config.ServiceConfig {
	ports := []string{}
	for _, p := range tainer.Ports {
		ports = append(ports, fmt.Sprintf("%d", p.Port))
	}

	return &config.ServiceConfig{
		Image:         tainer.Image,
		ContainerName: tainer.Name,
		Entrypoint:    tainer.Command,
		Ports:         ports,
	}
}

func (b *ComposeBackend) Start() error {
}

func (b *ComposeBackend) Stop() error {
}
