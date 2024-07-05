package main

import (
	"fmt"
	"github.com/openhue/openhue-go"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {

	bridge, key := loadConf()

	home, err := openhue.NewHome(bridge, key)
	if err != nil {
		log.Fatal(err)
	}

	lights, err := home.GetLights()
	if err != nil {
		return
	}

	for id, light := range lights {
		fmt.Printf("Toggling light %s (%s)\n", *light.Metadata.Name, id)
		home.UpdateLight(*light.Id, openhue.LightPut{
			On: light.Toggle(),
		})
	}
}

// loadConf lookup your Hue Bridge IP and Api Key from the well-known OpenHue standard configuration file.
func loadConf() (string, string) {

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
