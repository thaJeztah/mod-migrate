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
	fmt.Println(p.Hello())

	// FIXME(thaJeztah): we can't use a v1 plugin in v2 because the signature is different (v1 "pkg/foo" !== v2 "pkg/foo")
	// # github.com/thaJeztah/mod-migrate/v2
	//  ./main.go:18:18: cannot use p (variable of type *plugin.Plugin) as Plugin value in argument to PrintPluginName: *plugin.Plugin does not implement Plugin (wrong type for method Hello)
	// 		have Hello() *"github.com/thaJeztah/mod-migrate/pkg/foo".Hello
	// 		want Hello() *"github.com/thaJeztah/mod-migrate/v2/pkg/foo".Hello
	// PrintPluginName(p)
}

func PrintPluginName(p Plugin) {
	fmt.Println(p.Hello())
}
