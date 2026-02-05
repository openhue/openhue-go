package main

import (
	"context"
	"fmt"

	"github.com/openhue/openhue-go"
)

func main() {

	home, err := openhue.NewHome(openhue.LoadConfNoError())
	openhue.CheckErr(err)

	ctx := context.Background()
	lights, err := home.GetLights(ctx)
	openhue.CheckErr(err)

	for id, light := range lights {

		fmt.Printf("> Found light %s (%s)\n", *light.Metadata.Name, id)
	}
}
