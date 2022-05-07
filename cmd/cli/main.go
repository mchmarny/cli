package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mchmarny/twcli/pkg/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/urfave/cli"
)

const (
	name = "cli"
)

var (
	version = "v0.0.1-default"

	cfg *config.AppConfig
)

func main() {
	initLogging(name, version)

	homeDir, created, err := config.GetOrCreateHomeDir(name)
	fatalErr(err)
	log.Debug().Msgf("home dir (created: %v): %s", created, homeDir)

	cfg, err = config.ReadOrCreate(homeDir)

	app := &cli.App{
		Name:     "twee",
		Version:  fmt.Sprintf("%s", version),
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

func fatalErr(err error) {
	if err != nil {
		log.Fatal().Err(err).Msg("fatal error")
	}
}

func initLogging(name, version string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "ts"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"
	zerolog.ErrorFieldName = "err"
	zerolog.CallerFieldName = "caller"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
