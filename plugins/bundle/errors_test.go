package bundle

import (
	"fmt"
	"testing"
)

func TestErrors(t *testing.T) {
	errs := Errors{
		"foo": fmt.Errorf("foo error"),
		"bar": fmt.Errorf("bar error"),
	}

	expected := "2 error(s) occurred:\n'foo': foo error\n'bar': bar error"
	result := errs.Error()

	if result != expected {
		t.Errorf("Expected: %v \nbut got: %v", expected, result)
	}

	expected = "bar error"
	result = errs["bar"].Error()

	if result != expected {
		t.Errorf("Expected: %v \nbut got: %v", expected, result)
	}
}
