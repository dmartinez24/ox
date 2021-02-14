package standard

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Tester struct{}

func (t Tester) Name() string {
	return "standard/tester"
}

func (b *Tester) RunBeforeTest(ctx context.Context, root string, args []string) error {
	return os.Setenv("GO_ENV", "test")
}

func (p *Tester) Test(ctx context.Context, root string, args []string) error {
	fmt.Println(">>> Running Tests")

	cmd := exec.CommandContext(ctx, "go", p.testArgs(args)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running: %v\n", strings.Join(cmd.Args, " "))

	return cmd.Run()
}

func (p *Tester) testArgs(args []string) []string {
	base := []string{
		"test",
	}

	if !strings.Contains(strings.Join(args, " "), "-p") {
		base = append(base, "-p", "1")
	}

	cargs := append(base, "./...")
	if len(args) > 0 {
		cargs = append(base, args...)
	}

	return cargs
}
