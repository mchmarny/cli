package main

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

var (
	simpleCmd = cli.Command{
		Name:    "simple",
		Aliases: []string{"s"},
		Usage:   "Simple CLI command.",
		Subcommands: []cli.Command{
			{
				Name:    "one",
				Aliases: []string{"o"},
				Usage:   "First command.",
				Action:  cmdImplementation,
				Flags: []cli.Flag{
					stringFlag,
				},
			},
			{
				Name:    "two",
				Aliases: []string{"t"},
				Action:  cmdImplementation,
				Usage:   "Second command.",
				Flags: []cli.Flag{
					intFlag,
				},
			},
		},
	}
)

func cmdImplementation(c *cli.Context) error {
	val := c.String(stringFlag.Name)

	limit := c.Int(intFlag.Name)
	if limit == 0 {
		limit = limitDefault
	}

	log.Debug().Msgf("value: %s and %d", val, limit)

	list := []string{"one", "two", "three"}

	if err := json.NewEncoder(os.Stdout).Encode(list); err != nil {
		return errors.Wrapf(err, "error encoding list: %v", list)
	}

	return nil
}
