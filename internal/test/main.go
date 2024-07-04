package main

import (
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

	for _, light := range lights {

		if *light.Metadata.Name == "Bureau" {
			body := openhue.UpdateLightJSONRequestBody{
				On: light.On.Switch(),
			}
			home.SetLight(*light.Id, body)
		}
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
