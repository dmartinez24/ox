package version

import (
	"context"
	"fmt"

	"github.com/wawandco/oxpecker/plugins"
)

var (
	// The version of the CLI
	version = "0.0.1"
)

var (
	// Version is a Command
	_ plugins.Command = (*Command)(nil)
)

// Command command will print X version.
type Command struct{}

func (b Command) Name() string {
	return "version"
}

func (b Command) ParentName() string {
	return ""
}

func (b Command) HelpText() string {
	return "returns the current version of Oxpecker CLI"
}

// Run prints the version of the Oxpecker cli
func (b *Command) Run(ctx context.Context, root string, args []string) error {
	fmt.Printf("Oxpecker version %v\n", version)

	return nil
}
