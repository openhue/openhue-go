package openhue

// Toggleable defines resources that have an On field and can therefore be switched to on or off, mainly lights.
type Toggleable interface {
	Toggle() *On
	IsOn() bool
}
