package slices

import (
	"testing"
	"reflect"
)

func TestCopy(t *testing.T) {
	x := []string{"Hello", "World"}
	y := make([]string, 2, 2)

	// Values should not be equal
	if reflect.DeepEqual(x, y) {
		t.Fail()
	}

	// Copy the values
	Copy(&x, &y)

	// Values should be equal
	if !reflect.DeepEqual(x, y) {
		t.Fail()
	}
}

func TestAppend(t *testing.T) {
	x := []string{"Hello", "World"}
	Append(&x, "!")
	if !reflect.DeepEqual(x, []string{"Hello", "World", "!"}) {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	x := []string {"Hello", "World", "!"}
	Delete(&x, 1)
	if !reflect.DeepEqual(x, []string{"Hello", "!"}) {
		t.Fail()
	}
}
