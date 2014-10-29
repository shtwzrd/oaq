package oaq

import( "testing"
)


type TestComponent struct {
	BaseComponent
}

type AnotherComponent struct {
	BaseComponent
}

func TestAddComponent(t *testing.T) {
	c := new(TestComponent)
	entity := NewEntity()
	
	err := entity.Add(c)
	if err != nil {
		t.Error("Adding a Component to an Entity failed, even though the Entity " +
			"did not have a Component of that Type")
	} else {
		t.Log("TestAddComponent: PASS")
	}
	
}

func TestAddEntity(t *testing.T) {
	entity := NewEntity()
	subentity := NewNamedEntity("test")
	
	err := entity.Add(subentity)
	
	if err != nil {
		t.Error("Failed to add an Entity as a Component")
	} else {
		t.Log("TestAddEntity: PASS")
	}
}

func TestAddUnnamedEntity(t *testing.T) {
	entity := NewEntity()
	unnamedentity := NewEntity()
	
	err := entity.Add(unnamedentity)
	
	if err == nil {
		t.Error("Adding an unnamed Entity should fail, but it succeeded")
	} else {
		t.Log("TestAddUnnamedEntity: PASS")
	}
}

func TestAddExistingEntity(t *testing.T) {
	entity := NewEntity()
	subentityA := NewNamedEntity("turnip")
	subentityB := NewNamedEntity("turnip")
	
	entity.Add(subentityA)
	
	_, present := entity.components[subentityA.Name];
	if !present {
		t.Error("Entity does not store Entities by name")
	}
	
	err := entity.Add(subentityB)
	
	
	if err == nil {
		t.Error(`Adding an Entity with the same name as an already added
			one should fail, but it succeeded`)
	} else {
		t.Log("TestAddExistingEntity: PASS")
	}
}

func TestAddExistingComponent(t *testing.T) {
	c := new(TestComponent)
	d := new(TestComponent)
	z := new(AnotherComponent)
	entity := NewEntity()

	entity.Add(c)
	denied := entity.Add(z)
	failed := entity.Add(d)
	if failed == nil {
		t.Error("Adding a Component to an Entity succeeded, even though the Entity " +
			"already had a Component of that Type")
	} else if denied != nil {
		t.Error("Adding a Component to an Entity failed, even though the Entity " +
			"did not have a Component of that Type")
	} else {
		t.Log("TestAddExistingComponent: PASS")
	}
}

func TestRemoveComponent(t *testing.T) {
	c := new(TestComponent)
	entity := NewEntity()

	entity.Add(c)
	err := entity.Remove(c)
	if err != nil {
		t.Error("Removing a Component from an Entity failed, even though the Entity " +
			"had a Component of that Type")
	} else {
		t.Log("TestRemoveComponent: PASS")
	}
}

func TestRemoveNonExistingComponent(t *testing.T) {
	c := new(TestComponent)
	entity := NewEntity()

	failed := entity.Remove(c)
	if failed == nil {
		t.Error("Removing a Component from an Entity succeeded, even though the Entity " +
			"did not have Component of that Type")
	} else {
		t.Log("TestRemoveNonExistingComponent: PASS")
	}
}

func TestGetEntityOfRootEntity(t *testing.T) {
	root := NewEntity()
	c := new(TestComponent)
	d := new(AnotherComponent)

	root.Add(c)
	root.Add(d)

	child := NewEntity()
	x := new(TestComponent)
	y := new(AnotherComponent)
	child.Add(x)
	child.Add(y)
	root.Add(child)

	_, shouldError := root.Entity()
	if shouldError == nil {
		t.Error("Asking for the Entity of a Root Entity (an Entity which is not a" +
			" component of any other Entity) should result in an error.")
	} else {
		t.Log("TestGetEntityOfRootEntity: PASS")
	}
}

func TestEntityIDPersists(t *testing.T) {
	entity := NewEntity()

	subentity := NewNamedEntity("sub")
	entity.Add(subentity)
	_, ok := FindComponent(subentity.Id())
	if ok != nil {
		t.Error(`The Id on this side doesn't match the one in the
componentregistrar`)
	} else {
		t.Log(`TestEntityIDPersists: PASS`)
	}
}
func TestComponentIDPersists(t *testing.T) {
	entity := NewEntity()

	component := new(TestComponent)
	entity.Add(component)
	_, ok := FindComponent(component.Id())
	if ok != nil {
		t.Error(`The Id on this side doesn't match the one in the
componentregistrar`)
	} else {
		t.Log(`TestComponentIDPersists: PASS`)
	}
}

func TestFindComponent(t *testing.T) {
	c := new(TestComponent)
	d := new(AnotherComponent)
	entity := NewEntity()
	
	entity.Add(c)
	entity.Add(d)
	entity.Remove(d)
	
	_, wontError := FindComponent(c.Id())
	_, willError := FindComponent(d.Id())
	
	if wontError != nil {
		t.Error(`ComponentRegistrar cannot find the component, even
                though it's still attached to the Entity`)
	}
	if willError == nil {
		t.Error(`ComponenetRegistrar can still find the component, even
                though it has been removed from the entity`)
	}
	if wontError == nil && willError != nil {
		t.Log("TestAddExistingComponent: PASS")
	}
}
