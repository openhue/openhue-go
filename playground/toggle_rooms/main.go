// It is possible to modify all the lights of a room at once by accessing its GroupedLight service.
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
	rooms, err := home.GetRooms(ctx)
	openhue.CheckErr(err)

	for id, room := range rooms {

		fmt.Printf("> Toggling room %s (%s)\n", *room.Metadata.Name, id)

		for serviceId, serviceType := range room.GetServices() {

			if serviceType == openhue.ResourceIdentifierRtypeGroupedLight {
				groupedLight, _ := home.GetGroupedLightById(ctx, serviceId)

				home.UpdateGroupedLight(ctx, *groupedLight.Id, openhue.GroupedLightPut{
					On: groupedLight.Toggle(),
				})
			}
		}
	}
}
