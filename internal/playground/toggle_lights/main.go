package main

import (
	"fmt"
	"github.com/openhue/openhue-go"
	"log"
)

func main() {

	home, err := openhue.NewHome(openhue.LoadConf())
	if err != nil {
		log.Fatal(err)
	}

	lights, err := home.GetLights()
	if err != nil {
		return
	}

	for id, light := range lights {

		fmt.Printf("> Toggling light %s (%s)\n", *light.Metadata.Name, id)

		home.UpdateLight(*light.Id, openhue.LightPut{
			On: light.Toggle(),
		})
	}
}
