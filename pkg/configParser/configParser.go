package configParser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
	"olympos.io/encoding/edn"
)

func ReadFile(path string, cfg interface{}) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return ReadString(string(b), cfg)
}

func ReadString(serializedConfig string, cfg interface{}) error {

	err := yaml.Unmarshal([]byte(serializedConfig), cfg)
	if err == nil {
		return nil
	}
	err = json.Unmarshal([]byte(serializedConfig), cfg)
	if err == nil {
		return nil
	}
	err = toml.Unmarshal([]byte(serializedConfig), cfg)
	if err == nil {
		return nil
	}
	err = edn.Unmarshal([]byte(serializedConfig), cfg)
	if err == nil {
		return nil
	}

	return fmt.Errorf("config string could not be parsed")
}
