package help

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/wawandco/ox/plugins/core"
)

// printSingle prints help details for a passed plugin
// Usage, Subcommands and Flags.
func (h *Command) printSingle(command core.Command, names []string) {
	fmt.Println("Description:")
	if th, ok := command.(core.HelpTexter); ok {
		fmt.Printf("  %v\n\n", th.HelpText())
	}

	fmt.Println("Usage:")
	usage := fmt.Sprintf("  ox %v \n", command.Name())

	if command.ParentName() != "" {
		usage = fmt.Sprintf("  ox %v \n", strings.Join(names, " "))
	}

	th, isSubcommander := command.(core.Subcommander)
	if isSubcommander {
		usage = fmt.Sprintf("  ox %v [subcommand]\n", command.Name())
	}

	fmt.Println(usage)

	if isSubcommander {
		w := new(tabwriter.Writer)
		defer w.Flush()

		w.Init(os.Stdout, 8, 8, 3, '\t', 0)
		fmt.Println("Subcommands:")

		for _, scomm := range th.Subcommands() {
			if scomm.ParentName() == "" {
				continue
			}

			helpText := ""
			if ht, ok := scomm.(core.HelpTexter); ok {
				helpText = ht.HelpText()
			}

			fmt.Fprintf(w, "  %v\t%v\n", scomm.Name(), helpText)
		}
	}
	if th, ok := command.(core.Aliaser); ok {
		fmt.Printf("Alias: \n  %s\n", th.Alias())
	}

	if th, ok := command.(core.FlagParser); ok {
		fmt.Println("Flags:")

		flags := th.Flags()
		flags.SetOutput(os.Stderr)
		flags.PrintDefaults()
		fmt.Println("")

		return
	}

}
