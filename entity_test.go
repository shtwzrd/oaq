package oaq

import "testing"

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
