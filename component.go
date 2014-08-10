/*
Package oaq implements a simple library that can form the basis of an
Entity-Component System (ECS).
*/
package oaq

import (
	"errors"

	uuid "github.com/nu7hatch/gouuid"
)

// BaseComponent is a small struct that fulfills all the necessary functionality
// to implement the Component interface. BaseComponent is meant to be an
// anonymous embedded struct in every Component you create, as it frees
// you from having to put any wiring or clerical logic in the struct.
type BaseComponent struct {
	entity      *Entity
	subscribers []chan uuid.UUID
	id          *uuid.UUID
}

func (bc BaseComponent) Entity() (en *Entity, err error) {
	if bc.entity == nil {
		err = errors.New(`Component has no Entity; it is itself either a root
		Entity or a dangling Component.`)
	}
	return
}

func (bc *BaseComponent) setEntity(en *Entity) (err error) {
	if en == nil {
		err = errors.New("Provided Entity was nil!")
		return
	}
	// The Id of a component is automatically generated upon association
	// with an Entity, allowing us to avoid needing any explicit ctor function
	// for a component -- they're just structs with Plain-Old-Data.

	//Entities assign their own id on creation, and that shouldn't be reassigned
	if bc.id == nil { //If not nil, this component must be an Entity itself
		bc.id, err = uuid.NewV4()
		bc.entity = en
		registerComponent(bc)
	}
	return
}

// Subscribe allows a Processor to express its interest in a Component and
// receive a channel the Component can use to send its UUID to the Processor
// whenever Notify is fired. Subscribe is lazy-initialized -- no slice of
// channels is created until a Processor has subscribed to the Component.
func (bc *BaseComponent) Subscribe() chan uuid.UUID {
	if bc.subscribers == nil {
		bc.subscribers = make([]chan uuid.UUID, 0)
	}

	//Add our new subscriber
	var ch chan uuid.UUID
	bc.subscribers = append(bc.subscribers, ch)
	//And initialize the channel for our subscriber
	bc.subscribers[len(bc.subscribers)-1] = make(chan uuid.UUID)
	return bc.subscribers[len(bc.subscribers)-1]
}

// Notify should be called whenever a Processor modifies a Component. It is the
// responsibility of the Processor to call Notify in this event, to alert any
// interested Processors that the data of the Component has been invalidated.
// In the event that no Processor is interested in the Component, Notify does
// nothing.
func (bc BaseComponent) Notify() {
	if bc.subscribers != nil {
		for _, ch := range bc.subscribers {
			ch <- *bc.id
		}
	}
}

func (bc BaseComponent) Id() uuid.UUID {
	return *bc.id
}

type Component interface {
	Id() uuid.UUID
	Entity() (*Entity, error)
	setEntity(*Entity) error
	Notify()
	Subscribe() chan uuid.UUID
}
