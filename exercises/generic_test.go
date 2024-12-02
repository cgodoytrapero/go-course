package exercises

import (
	"testing"
)

func TestGivenAnyOfWhenIntegerSimpleComparisonThenTrue(t *testing.T) {
	if AnyOf([]int{1, 2, 3},
		func(i int) bool { return i == 2 }) == false {
		t.Errorf("No works")
	}
}

func TestGivenAnyOfWhenSpringComparisonThenTrue(t *testing.T) {
	if AnyOf([]string{"foo", "bar", "baz"},
		func(i string) bool { return i == "bar" }) == false {
		t.Errorf("No works")
	}
}
