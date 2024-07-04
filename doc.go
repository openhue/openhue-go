/*
Package openhue provides a simple API on top of the Philips Hue CLIP API. Most of the code is automatically generated
thanks to https://github.com/openhue/openhue-api, the main project of the OpenHue organization.

The main concept of this library is the Home abstraction that acts as an entry point to the rest of the resources exposed
by the Philips Hue bridge.

Let's start creating your Home and listing all the rooms:

	func main() {
		h, _ := openhue.NewHome("192.168.0.1", "replace with your actual API key")
		rooms, _ := home.GetRooms()
		for _, room := range rooms {
			fmt.Println(*room.Metadata.Name)
		}
	}

Here we create a new Home instance from a given Bridge IP and API key. Then we are able to list all rooms of type RoomGet
that are contained in your Home. Please note that we explicitly ignored all the errors to simply this snippet.
*/
package openhue
