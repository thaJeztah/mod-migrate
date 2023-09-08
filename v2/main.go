package main

import (
	"fmt"

	"github.com/thaJeztah/mod-migrate/plugin"
	"github.com/thaJeztah/mod-migrate/v2/pkg/foo"
)

type Plugin interface {
	Hello() *foo.Hello
}

func main() {
	p := plugin.NewPlugin("v1.0.0 plugin")
	PrintPluginName(p)
}

func PrintPluginName(p Plugin) {
	fmt.Println(p.Hello())
}
