package main

import (
	"fmt"
	"github.com/openhue/openhue-go"
)

func main() {

	home, err := openhue.NewHome(openhue.LoadConf())
	openhue.CheckErr(err)

	lights, err := home.GetLights()
	openhue.CheckErr(err)

	for id, light := range lights {

		fmt.Printf("> Toggling light %s (%s)\n", *light.Metadata.Name, id)

		home.UpdateLight(*light.Id, openhue.LightPut{
			On: light.Toggle(),
		})
	}
}
