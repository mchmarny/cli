package main

import (
	"os"
	"testing"

	"github.com/mchmarny/cli/pkg/config"
	"github.com/rs/zerolog/log"
)

const (
	testDir = "../../tmp"
)

func TestMain(m *testing.M) {
	os.RemoveAll(testDir)
	initLogging()

	cfg, err := config.ReadOrCreate(testDir)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read or create config")
	}
	log.Debug().Msgf("config: %+v", cfg)

	code := m.Run()
	os.Exit(code)
}
