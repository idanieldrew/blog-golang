package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func Read(path string, cfg *Config) error {
	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		return parseYaml(path, cfg)
	default:
		return errors.New("format is not correct")
	}
}

func parseYaml(path string, cfg *Config) error {
	// open path
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		if ce := file.Close(); ce != nil {
			err = ce
		}
	}()

	// return yaml error(ey)
	if ye := yaml.NewDecoder(file).Decode(cfg); ye != nil {
		return ye
	}

	return nil
}
