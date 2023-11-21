package configParser

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadFile(path string, cfg interface{}) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return ReadString(b, cfg)
}

func ReadString(buf []byte, cfg interface{}) error {
	err := yaml.Unmarshal(buf, cfg)
	return err
}
