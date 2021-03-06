// +build standalone

package main

import (
	"github.com/tonistiigi/buildkit_poc/control"
	"github.com/urfave/cli"
)

func appendFlags(f []cli.Flag) []cli.Flag {
	return f
}

// root must be an absolute path
func newController(c *cli.Context, root string) (*control.Controller, error) {
	return control.NewStandalone(root)
}
