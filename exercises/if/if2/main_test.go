// if2
// Make me compile!

package main_test

import "testing"

func fooIfFizz(fizzish string) string {
	switch fizzish {
	case "fizz":
		return "foo"
	case "fuzz":
		return "bar"
	default:
		return "baz"
	}
}

func TestFooForFizz(t *testing.T) {
	result := fooIfFizz("fizz")
	if result != "foo" {
		t.Errorf("should be 'foo' but got %s", result)
	}
}

func TestBarForFuzz(t *testing.T) {
	result := fooIfFizz("fuzz")
	if result != "bar" {
		t.Errorf("should be 'bar' but got %s", result)
	}
}

func TestDefaultForBazz(t *testing.T) {
	result := fooIfFizz("random stuff")
	if result != "baz" {
		t.Errorf("should be 'baz' but got %s", result)
	}
}
