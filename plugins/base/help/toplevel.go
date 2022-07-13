package help

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/wawandco/ox/plugins/core"
)

// printTopLevel prints the top level help text with a table that contains top level
// commands (names) and descriptions.
func (c *Command) printTopLevel() {
	fmt.Println("Usage")
	fmt.Printf("  ox [command]\n\n")

	w := new(tabwriter.Writer)
	defer w.Flush()

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 3, '\t', 0)
	fmt.Print("Available top level Commands:\n\n")
	fmt.Println("Command\t     Alias")

	for _, plugin := range c.commands {
		helpText := ""
		if ht, ok := plugin.(core.HelpTexter); ok {
			helpText = ht.HelpText()
		}

		if p, ok := plugin.(core.Aliaser); ok {
			fmt.Fprintf(w, "  %v\t%v\t%v\n", plugin.Name(), p.Alias(), helpText)
		} else {
			fmt.Fprintf(w, "  %v\t\t%v\n", plugin.Name(), helpText)
		}
	}

}
