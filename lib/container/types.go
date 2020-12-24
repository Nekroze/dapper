package container

type PortExposure struct {
	Port uint
	Scheme string
}

type Container struct {
	Name string
	Image string
	Ports PortExposure[]
	Command string[]
}

type Application map[string]Container

type Backend interface {
	Start() error
	Stop() error
}
