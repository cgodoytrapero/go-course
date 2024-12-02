package exercises

import (
	"testing"
)

func TestGivenContainerWhenApplyFlowThenOk(t *testing.T) {
	container := Container{}

	container.Add(1)
	container.Add(3)

	if container.Exist(1) == false {
		t.Errorf("container.Exist(1) = false; want true")
	}

	if container.Exist(2) == true {
		t.Errorf("container.Exist(2) = true; want false")
	}

	if container.Delete(1) == false {
		t.Errorf("container.Delete(1) = false; want true")
	}

	if container.Exist(1) == true {
		t.Errorf("container.Exist(1) = true; want false")
	}
}
