package main

import (
	"gopkg.in/urfave/cli.v2"
	"log"
	"os"
)

const (
	flagDataPath   = "dataPath"
	flagDataValues = "values"
	flagGoImports  = "enableGoImports"
)

func main() {
	app := cli.App{
		Name:    "genius",
		Version: "v1.0.0",
		Usage:   "various generation commands, refer to each command help for usage",
		Commands: []*cli.Command{
			Tmpl,
		},
		EnableShellCompletion: true,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
