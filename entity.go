package oaq

import (
	"errors"
	"fmt"
	"reflect"
)

type Entity struct {
	BaseComponent
	components map[string]Component
	Name       string
}

func NewEntity() *Entity {
	entity := new(Entity)
	entity.components = make(map[string]Component)
	var err error
	//	entity.id, err = uuid.NewV4()
	if err != nil {
		fmt.Println("UUID error: ", err)
	}
	return entity
}

func NewNamedEntity(name string) (en *Entity) {
	en = NewEntity()
	en.Name = name
	return
}

// Adds a Component to an Entity. Returns an error if the Entity already had
// a Component of the same type.
func (en *Entity) Add(c Component) (err error) {
	t := reflect.TypeOf(c)

	if t == reflect.TypeOf(en) {
		ent := c.(*Entity)
		if ent.Name == "" {
			errmsg := fmt.Sprintf(`Adding an Entity with no name is
			not a supported operation, and Entity with reference %v
                        has no name`, &ent)
			err = errors.New(errmsg)
		} else {
			_, present := en.components[ent.Name]

			if present {
				errmsg := fmt.Sprintf(`Entity with reference %v, name %s is already assigned an
			entity with name %s`, &en, en.Name, t.String())
				err = errors.New(errmsg)
			} else {
				en.components[ent.Name] = ent
				en.registerComponent(c)
			}
		}

	} else {
		_, present := en.components[t.String()]

		if present {
			errmsg := fmt.Sprintf(`Entity with reference %v, name %s is already assigned a  
			component of type %s`, &en, en.Name, t.String())
			err = errors.New(errmsg)
		} else {
			en.registerComponent(c)
			en.components[t.String()] = c
			fmt.Println(c.Id())
		}
	}
	
	if err == nil {
		en.Notify() // Let interested Processors know we changed
	}
	return
}

// Removes a Component from an Entity. Returns an error if attempting to Remove
// a Component that the Entity did not have.
func (en *Entity) Remove(c Component) (err error) {
	t := reflect.TypeOf(c)
	if t == reflect.TypeOf(en) {
		ent := c.(*Entity)
		_, present := en.components[ent.Name]

		if !present {
			errmsg := fmt.Sprintf(`Entity with reference %v, name %s
			has no entity with name %s`, &en, en.Name, ent.Name)
			err = errors.New(errmsg)
		} else {
			err = unregisterComponent(ent)
			delete(en.components, ent.Name)
		}
	} else {
		_, present := en.components[t.String()]

		if present {
			fmt.Println("Was present in Entity")
			fmt.Println(c.Id())
			err = unregisterComponent(c)
			delete(en.components, t.String())
		} else {
			errmsg := fmt.Sprintf(`Entity with reference %v, name %s
                        has no component of type %v`, &en, en.Name, t)
			err = errors.New(errmsg)
		}
	}

	if err == nil {
		en.Notify() // Is anybody out there? I lost a component!
	}
	return
}
