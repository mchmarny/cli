package main

import (
	"github.com/urfave/cli"
)

const (
	limitDefault = 100
)

var (
	stringFlag = &cli.StringFlag{
		Name:     "username",
		Usage:    "Twitter username",
		Required: true,
	}

	intFlag = &cli.IntFlag{
		Name:  "limit",
		Usage: "Limit the number of results",
		Value: limitDefault,
	}
)
