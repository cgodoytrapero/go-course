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

func TestGivenAnyOfWhenStringComparisonThenTrue(t *testing.T) {
	if AnyOf([]string{"foo", "bar", "baz"},
		func(i string) bool { return i == "bar" }) == false {
		t.Errorf("No works")
	}
}

func TestGivenFindIfWhenFoundStringComparisonThenOne(t *testing.T) {
	if FindIf([]string{"foo", "bar", "baz"},
		func(i string) bool { return i == "bar" }) != 1 {
		t.Errorf("No works")
	}
}

func TestGivenAdjacentFindWhenNotFoundStringComparisonThenMinusOne(t *testing.T) {
	if AdjacentFind([]string{"foo", "bar", "baz"},
		func(i string, j string) bool { return i == j }) != -1 {
		t.Errorf("No works")
	}
}

func TestGivenAdjacentFindWhenFoundStringComparisonThenOne(t *testing.T) {
	if AdjacentFind([]string{"foo", "bar", "bar"},
		func(i string, j string) bool { return i == j }) != 1 {
		t.Errorf("No works")
	}
}

func TestGivenEqualWhenBothAreEqualThenTrue(t *testing.T) {
	if Equal([]string{"foo", "bar", "bar"},
		[]string{"foo", "bar", "bar"}) != true {
		t.Errorf("No works")
	}
}

func TestGivenEqualWhenBothAreNotEqualThenTrue(t *testing.T) {
	if Equal([]string{"foo", "bar", "bar"},
		[]string{"foo", "baz", "bar"}) != false {
		t.Errorf("No works")
	}
}

func TestGivenEqualWhenOneAreNilThenFalse(t *testing.T) {
	if Equal([]string{"foo", "bar", "bar"},
		nil) != false {
		t.Errorf("No works")
	}
}

func TestGivenReplaceIflWhenComplainsAndExistsThenOne(t *testing.T) {
	origin := []string{"foo", "bar", "bar"}
	if ReplaceIf(origin,
		"carlos",
		func(i string) bool { return i == "bar" }) != 1 {
		t.Errorf("No works")
	}

	if origin[1] != "carlos" {
		t.Errorf("No works")
	}
}

func TestGivenRemoveIflWhenComplainsAndExistsThenOne(t *testing.T) {
	origin := []string{"foo", "bar", "baz"}
	if !RemoveIf(&origin,
		func(i string) bool { return i != "bar" }) {
		t.Errorf("No works")
	}

	if len(origin) != 1 {
		t.Errorf("No works")
	}
}
