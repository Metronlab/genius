package main

import (
	"errors"
	"log"
	"os"
	"syscall"

	"github.com/Metronlab/genius/internal/geniuserr"
	"gopkg.in/urfave/cli.v2"
)

const flagDryRun = "dry"

func main() {
	app := cli.App{
		Name:    "genius",
		Version: "v1.0.0",
		Usage:   "various generation commands, refer to each command help for usage",
		Commands: []*cli.Command{
			Tmpl,
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagDryRun,
				EnvVars: []string{"GENIUS_DRY"},
				Usage: "Enable dry run to ensure resulting file is matching freshly generated one without modification. " +
					"Mismatch will result in error with status code 2",
				Value: false,
			},
		},
		EnableShellCompletion: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)

		if errors.Is(err, geniuserr.ErrDryMismatch) {
			syscall.Exit(2)
		}
		syscall.Exit(1)
	}
}
