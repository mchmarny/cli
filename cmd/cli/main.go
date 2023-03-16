package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mchmarny/cli/pkg/config"
	"github.com/mchmarny/cli/pkg/data"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

const (
	name           = "cli"
	limitDefault   = 100
	logLevelEnvVar = "debug"
)

var (
	version = "v0.0.1-default"
	commit  = "none"
	date    = "unknown"

	cfg *config.Config

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

func main() {
	initLogging()

	if version == "" || commit == "" || date == "" {
		fatalErr(errors.New("version, commit, and date must be set"))
	}

	compileTime, err := time.Parse("2006-01-02T15:04:05Z", date)
	if err != nil {
		log.Debug().Msg("compile time not set, using current")
		compileTime = time.Now()
	}
	dateStr := compileTime.UTC().Format("2006-01-02 15:04 UTC")

	homeDir, created, err := config.GetOrCreateHomeDir(name)
	fatalErr(err)
	log.Debug().Msgf("home dir (created: %v): %s", created, homeDir)

	cfg, err = config.ReadOrCreate(homeDir)
	fatalErr(err)

	if err = data.Init(homeDir); err != nil {
		fatalErr(err)
	}
	defer data.Close()

	app := &cli.App{
		Name:     "twee",
		Version:  fmt.Sprintf("%s (commit: %s, built: %s, config: %s)", version, commit, dateStr, cfg.Value),
		Compiled: time.Now(),
		Usage:    "cli",
		Commands: []cli.Command{
			simpleCmd,
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Print(err)
	}
}

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

func fatalErr(err error) {
	if err != nil {
		log.Fatal().Err(err).Msg("fatal error")
	}
}

func initLogging() {
	level := zerolog.InfoLevel
	levStr := os.Getenv(logLevelEnvVar)
	if levStr == "true" {
		level = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(level)

	out := zerolog.ConsoleWriter{
		Out: os.Stderr,
		PartsExclude: []string{
			zerolog.TimestampFieldName,
		},
	}

	log.Logger = zerolog.New(out)
}
