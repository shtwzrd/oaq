package oaq

import (
	"errors"

	uuid "github.com/nu7hatch/gouuid"
)

var (
	components map[uuid.UUID]Component
)

func registerComponent(c Component) (err error) {
	var empty [16]byte
	if c.Id() == empty {
		err = errors.New("Component was never given an id and cannot be registered!")
	}
	if components == nil {
		components = map[uuid.UUID]Component{}
	}
	components[c.Id()] = c
	return
}

func unregisterComponent(id uuid.UUID) (err error) {
	_, present := components[id]
	if !present {
		err = errors.New("Given id does not exist in the registry!")
	}
	delete(components, id)
	return
}

func FindComponent(id uuid.UUID) (c Component, err error) {
	c, present := components[id]

	if !present {
		err = errors.New("Given id does not exist in the registry!")
	}
	return
}
