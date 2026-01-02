package hashes

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Addr         string `env:"HASHES_ADDR"              envDefault:":8082"`
	CC           string `env:"HASHES_CC"                envDefault:"380"`
	NDCS         []int  `env:"HASHES_NDCS"              envSeparator:"," envDefault:"67"`
	NDCCap       int    `env:"HASHES_NDC_CAPACITY"      envDefault:"10000000"`
	Salt         string `env:"HASHES_SALT"              envDefault:"mySalt"`
	MsisdnLenMin int    `env:"HASHES_MSISDN_LENGTH_MIN" envDefault:"12"`
	MsisdnLenMax int    `env:"HASHES_MSISDN_LENGTH_MAX" envDefault:"21"`
}

func NewConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("❌ Failed to parse environment: %v", err)
	}
	if err := cfg.Validate(); err != nil {
		log.Fatalf("❌ Invalid configuration: %v", err)
	}
	return cfg
}

func (c *Config) Validate() error {
	if c.Addr == "" {
		return fmt.Errorf("HASHES_ADDR must not be empty")
	}
	if len(c.CC) < 2 {
		return fmt.Errorf("HASHES_CC must be a valid country code")
	}
	if len(c.NDCS) == 0 {
		return fmt.Errorf("HASHES_NDCS must contain at least one element")
	}
	if c.NDCCap <= 0 {
		return fmt.Errorf("HASHES_NDC_CAPACITY must be greater than zero")
	}
	if c.MsisdnLenMin <= 0 || c.MsisdnLenMax <= 0 || c.MsisdnLenMax < c.MsisdnLenMin {
		return fmt.Errorf("invalid MSISDN length range: min=%d, max=%d", c.MsisdnLenMin, c.MsisdnLenMax)
	}
	if strings.TrimSpace(c.Salt) == "" {
		return fmt.Errorf("HASHES_SALT must not be empty")
	}
	return nil
}

func LoadConfigFromEnv(envs map[string]string) (*Config, error) {
	for k, v := range envs {
		err := os.Setenv(k, v)
		if err != nil {
			return nil, err
		}
	}
	defer func() {
		for k := range envs {
			err := os.Unsetenv(k)
			if err != nil {
				return
			}
		}
	}()
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, cfg.Validate()
}
