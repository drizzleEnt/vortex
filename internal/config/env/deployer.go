package env

import (
	"errors"
	"os"
	"strconv"
	"time"
)

const (
	durationName = "DEPLOYER_DURATION"
)

type deployerConfig struct {
	dur time.Duration
}

func NewDeployerConfig() (*deployerConfig, error) {
	durationStr := os.Getenv(durationName)
	if len(durationStr) == 0 {
		return nil, errors.New("deployer duration not found")
	}

	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		return nil, err
	}

	return &deployerConfig{
		dur: time.Duration(duration),
	}, nil
}

func (cfg *deployerConfig) Duration() time.Duration {
	return cfg.dur
}
