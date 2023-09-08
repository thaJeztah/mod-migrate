package plugin

import "github.com/thaJeztah/mod-migrate/pkg/foo"

func NewPlugin(name string) *Plugin {
	return &Plugin{Name: name}
}

type Plugin struct {
	Name string
}

func (p *Plugin) Hello() *foo.Hello {
	return &foo.Hello{Name: p.Name}
}
