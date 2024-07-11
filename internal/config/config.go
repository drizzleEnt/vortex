package config

import (
	"time"

	"github.com/subosito/gotenv"
)

func Load(path string) error {
	err := gotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

type HTTPConfig interface {
	Address() string
}

type PGConfig interface {
	Address() string
}

type DeployerConfig interface {
	Duration() time.Duration
}
