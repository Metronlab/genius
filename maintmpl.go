package main

import (
	"github.com/Metronlab/genius/internal/geniuscmd/tmpl"
	"github.com/Metronlab/genius/internal/geniusio"
	"github.com/Metronlab/genius/internal/geniustypes"
	"gopkg.in/urfave/cli.v2"
)

const (
	flagDataPath   = "dataPath"
	flagDataFormat = "dataFormat"
	flagDataValues = "values"
	flagGoImports  = "enableGoImports"
)

var Tmpl = &cli.Command{
	Name:  "tmpl",
	Usage: "allow usage of native go templating framework for go code itself",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    flagDataPath,
			Aliases: []string{"data", "d"},
			Usage: "path to your data file like `file.json` in anyone of supported serialisation format, " +
				"if not specified, use stdin as text entry",
		},
		&cli.StringFlag{
			Name:    flagDataFormat,
			Aliases: []string{"e"},
			Usage:   "specify input data format and overide extension, accept json, yaml, toml and text",
		},
		&cli.GenericFlag{
			Name:    flagDataValues,
			Aliases: []string{"v"},
			Usage:   "values given with the format `key=value` that will be added to your accessible data",
			Value:   make(geniustypes.ValuesMap),
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
			c.String(flagDataFormat),
			c.Generic(flagDataValues).(geniustypes.ValuesMap),
			c.Args().Slice(),
			c.Bool(flagGoImports),
			geniusio.GetGenerationWriteFunc(c.Bool(flagDryRun)),
		)
	},
}
