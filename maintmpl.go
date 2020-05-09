package main

import (
	"github.com/Metronlab/genius/internal/tmpl"
	"gopkg.in/urfave/cli.v2"
)

var Tmpl = &cli.Command{
	Name:  "tmpl",
	Usage: "allow usage of native go templating framework for go code itself",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    flagDataPath,
			Aliases: []string{"data", "d"},
			Usage:   "path to your data file like `file.json` in anyone of supported serialisation format",
		},
		&cli.GenericFlag{
			Name:    flagDataValues,
			Aliases: []string{"v"},
			Usage:   "values given with the format `key=value` that will be added to your accessible data",
			Value:   make(tmpl.ValuesList),
		},
		&cli.BoolFlag{
			Name:    flagGoImports,
			Aliases: []string{"i"},
			Usage:   "enable usage of go imports ",
			Value:   false,
		},
	},
	Action: func(c *cli.Context) error {
		return tmpl.Tmpl(
			c.String(flagDataPath),
			c.Generic(flagDataValues).(tmpl.ValuesList),
			c.Bool(flagGoImports),
			c.Args().Slice(),
		)
	},
}
