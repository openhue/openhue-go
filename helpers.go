package openhue

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Toggleable defines resources that have an On field and can therefore be switched to on or off, mainly lights.
type Toggleable interface {
	Toggle() *On
	IsOn() bool
}

// LoadConf looks up your Hue Bridge IP and Api Key from the well-known OpenHue standard configuration file.
func LoadConf() (string, string) {

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	yamlFile, err := os.ReadFile(homedir + "/.openhue/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	c := make(map[string]interface{})

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatal(err)
	}

	return c["bridge"].(string), c["key"].(string)
}
