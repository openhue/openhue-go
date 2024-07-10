package openhue

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// Toggleable defines resources that have an On field and can therefore be switched to on or off, mainly lights.
type Toggleable interface {
	Toggle() *On
	IsOn() bool
}

type Conf struct {
	bridgeIP string
	apiKey   string
}

// LoadConf looks up your Hue Bridge IP and Api Key from the well-known OpenHue standard configuration file.
func LoadConf() (*Conf, error) {

	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("unable to get home directory: %w", err)
	}

	yamlFile, err := os.ReadFile(homedir + "/.openhue/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("unable to read ~/.openhue/config.yaml: %w", err)
	}

	c := make(map[string]interface{})

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse ~/.openhue/config.yaml: %w", err)
	}

	return &Conf{c["bridge"].(string), c["key"].(string)}, nil
}

// LoadConfNoError is similar to LoadConf() except that it will fatal if there are any errors.
func LoadConfNoError() (string, string) {
	conf, err := LoadConf()
	CheckErr(err)
	return conf.bridgeIP, conf.apiKey
}

// CheckErr prints the msg with the prefix 'Error:' and exits with error code 1. If the msg is a nil, it does nothing.
func CheckErr(msg error) {
	if msg != nil {
		fmt.Fprintln(os.Stderr, "Error:", msg)
		os.Exit(1)
	}
}
