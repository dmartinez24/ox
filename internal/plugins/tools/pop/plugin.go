package pop

import (
	"github.com/paganotoni/x/internal/plugins"
)

var (
	_ plugins.Plugin = (*Plugin)(nil)
)

type Plugin struct {
	// subcommands we will invoke depending on parameters
	// these are filled when Receive is called.
	subcommands []plugins.Subcommand
}

func (p *Plugin) Name() string {
	return "pop"
}
