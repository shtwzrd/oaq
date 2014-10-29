package oaq

import (
	"errors"
	"fmt"

	uuid "github.com/nu7hatch/gouuid"
)

var (
	components map[uuid.UUID]Component
)

func (en *Entity) registerComponent(c Component) (err error) {
	if en == nil {
		err = errors.New("Provided Entity was nil!")
		return
	}

	var empty [16]byte
	
	c.init()
	if c.Id() == empty {
		c.init()
	}
	
	// The Id of a component is automatically generated upon association
	// with an Entity, allowing us to avoid needing any explicit ctor function
	// for a component -- they're just structs with Plain-Old-Data.
	
	//Entities assign their own id on creation, and that shouldn't be reassigned
	if en.id == nil  {
		en.init()
	}
	
	fmt.Println(en)
	fmt.Println(c.Id())
	c.setEntity(en)
	
	if components == nil {
		components = map[uuid.UUID]Component{}
	}
	if c.Id() == empty {
		c.init()
	}
	
	components[c.Id()] = c
	return
}

func unregisterComponent(c Component) (err error) {
	fmt.Printf("Trying... %v", c.Id())
	_, present := components[c.Id()]
	if !present {
		err = errors.New("Given id does not exist in the registry!")
		fmt.Println("Failed to find component")
	}
	delete(components, c.Id())
	return
}

func FindComponent(id uuid.UUID) (c Component, err error) {
	c, present := components[id]
	
	if !present {
		err = errors.New("Given id does not exist in the registry!")
	}
	return
}
