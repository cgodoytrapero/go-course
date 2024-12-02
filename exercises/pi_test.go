package exercises

import (
	"testing"
)

func TestGivenPiFunctionWhenNValueIs1Then10(t *testing.T) {
	result := Pi(1)
	expected := 3.132178
	if result != expected {
		t.Errorf("Pi(1) = %f; want %f", result, expected)
	}
}

func TestGivenPiFunctionWhenNValueIs10Then3140237(t *testing.T) {
	result := Pi(10)
	expected := 3.140237
	if result != expected {
		t.Errorf("Pi(1) = %f; want %f", result, expected)
	}
}
