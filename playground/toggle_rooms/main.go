// It is possible to modify all the lights of a room at once by accessing its GroupedLight service.
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

	rooms, err := home.GetRooms()
	if err != nil {
		return
	}

	for id, room := range rooms {

		fmt.Printf("> Toggling room %s (%s)\n", *room.Metadata.Name, id)

		for serviceId, serviceType := range room.GetServices() {

			if serviceType == openhue.ResourceIdentifierRtypeGroupedLight {
				groupedLight, _ := home.GetGroupedLightById(serviceId)

				home.UpdateGroupedLight(*groupedLight.Id, openhue.GroupedLightPut{
					On: groupedLight.Toggle(),
				})
			}
		}
	}
}
