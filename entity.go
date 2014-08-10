package oaq

import (
	"errors"
	"fmt"
	"reflect"

	uuid "github.com/nu7hatch/gouuid"
)

type Entity struct {
	BaseComponent
	components map[reflect.Type]Component
	name       string
}

func NewEntity() *Entity {
	entity := new(Entity)
	entity.components = make(map[reflect.Type]Component)
	var err error
	entity.id, err = uuid.NewV4()
	if err != nil {
		fmt.Println("UUID error: ", err)
	}
	registerComponent(entity)
	return entity
}

func NewNamedEntity(name string) (en *Entity) {
	en = NewEntity()
	en.name = name
	return
}

// Adds a Component to an Entity. Returns an error if the Entity already had
// a Component of the same type.
func (en *Entity) Add(c Component) (err error) {
	t := reflect.TypeOf(c)
	_, present := en.components[t]

	if present {
		errmsg := fmt.Sprintf(`Entity with reference %v is already assigned a  
			component of type %v`, &en, t)
		err = errors.New(errmsg)
	} else {
		en.components[t] = c
		c.setEntity(en) //Give the Component a reference to this Entity
	}
	en.Notify() // Let interested Processors know we changed
	return
}

// Removes a Component from an Entity. Returns an error if attempting to Remove
// a Component that the Entity did not have.
func (en *Entity) Remove(c Component) (err error) {
	t := reflect.TypeOf(c)
	_, present := en.components[t]

	if present {
		delete(en.components, t)
	} else {
		errmsg := fmt.Sprintf(`Entity with reference %v has no 
			component of type %v`, &en, t)
		err = errors.New(errmsg)
	}
	en.Notify() // Is anybody out there? I lost a component!
	return
}
