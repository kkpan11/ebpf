package internal

import (
	"testing"

	"github.com/go-quicktest/qt"
)

func TestIdentifier(t *testing.T) {
	testcases := []struct {
		in, out string
	}{
		{".rodata", "Rodata"},
		{"_foo_bar_", "FooBar"},
		{"ipv6_test", "Ipv6Test"},
		{"FOO_BAR", "FOO_BAR"},
		{"FOO_", "FOO_"},
		{"FOO__BAR", "FOO__BAR"},
		{"FOO___BAR", "FOO___BAR"},
		{"_FOO__BAR", "FOO__BAR"},
		{"__FOO__BAR", "FOO__BAR"},
	}

	for _, tc := range testcases {
		have := Identifier(tc.in)
		if have != tc.out {
			t.Errorf("Expected %q as output of %q, got %q", tc.out, tc.in, have)
		}
	}
}

type foo struct{}

func TestGoTypeName(t *testing.T) {
	type bar[T any] struct{}
	qt.Assert(t, qt.Equals(GoTypeName(foo{}), "foo"))
	qt.Assert(t, qt.Equals(GoTypeName(new(foo)), "foo"))
	qt.Assert(t, qt.Equals(GoTypeName(new(*foo)), "foo"))
	qt.Assert(t, qt.Equals(GoTypeName(bar[int]{}), "bar[int]"))
	qt.Assert(t, qt.Equals(GoTypeName(bar[foo]{}), "bar[foo]"))
	qt.Assert(t, qt.Equals(GoTypeName(bar[testing.T]{}), "bar[testing.T]"))
}
