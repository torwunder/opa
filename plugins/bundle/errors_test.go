package bundle

import (
	"errors"
	"fmt"
	"github.com/open-policy-agent/opa/download"
	"testing"
)

func TestErrors(t *testing.T) {
	errs := Errors{
		Error{name: "foo", err: fmt.Errorf("foo error")},
		Error{name: "bar", err: fmt.Errorf("bar error")},
	}

	expected := "'foo': foo error\n'bar': bar error"
	result := errs.Error()

	if result != expected {
		t.Errorf("Expected: %v \nbut got: %v", expected, result)
	}
}

func TestUnwrapSlice(t *testing.T) {
	fooErr := Error{name: "foo", err: fmt.Errorf("foo error")}
	barErr := Error{name: "bar", err: fmt.Errorf("bar error")}

	errs := Errors{fooErr, barErr}

	result := errs.Unwrap()

	if result[0].Error() != fooErr.Error() {
		t.Fatalf("expected %v \nbut got: %v", fooErr, result[0])
	}
	if result[1].Error() != barErr.Error() {
		t.Fatalf("expected %v \nbut got: %v", barErr, result[1])
	}
}

func TestUnwrap(t *testing.T) {
	fooErr := Error{name: "foo", err: download.HTTPError{StatusCode: 500}}
	barErr := Error{name: "bar", err: download.HTTPError{StatusCode: 400}}

	errs := Errors{fooErr, barErr}

	// unwrap first bundle.Error
	var bundleError Error
	if !errors.As(errs, &bundleError) {
		t.Fatal("failed to unwrap Error")
	}
	if bundleError.Error() != fooErr.Error() {
		t.Fatalf("expected: %v \ngot: %v", fooErr, bundleError)
	}

	// unwrap first HTTPError
	var httpError download.HTTPError
	if !errors.As(errs, &httpError) {
		t.Fatal("failed to unwrap Error")
	}
	if httpError.Error() != fooErr.err.Error() {
		t.Fatalf("expected: %v \ngot: %v", fooErr.err, httpError)
	}

	// unwrap HTTPError from bundle.Error
	if !errors.As(bundleError, &httpError) {
		t.Fatal("failed to unwrap HTTPError")
	}
	if httpError.Error() != fooErr.err.Error() {
		t.Fatalf("expected: %v \nbgot: %v", fooErr.err, httpError)
	}
}
