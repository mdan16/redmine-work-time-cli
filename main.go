package main

import (
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v2"
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	var file string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Destination: &file,
				Required:    true,
			},
		},
		Name:  "redmine-work-time-cli",
		Usage: "",
		Action: func(cCtx *cli.Context) error {
			cmd := &CMD{
				outStream: c.outStream,
				errStream: c.errStream,
			}
			return cmd.Run(file)
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprint(c.errStream, err)
		return 1
	}
	return 0
}

func main() {
	cli := &CLI{
		outStream: os.Stdout,
		errStream: os.Stderr,
	}
	os.Exit(cli.Run(os.Args))
}
