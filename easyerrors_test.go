package easyerrors

import "testing"
import "github.com/stretchr/testify/require"

//goland:noinspection GoPrintFunctions
func TestAs(t *testing.T) {
	type fooable interface {
		Foo()
	}

	err := Errorf("bar: %w", errFoo{})

	var errHello fooable
	require.True(t, As(err, &errHello))
	require.False(t, As(New("bar"), &errHello))
}

//goland:noinspection GoDfaNilDereference,GoPrintFunctions
func TestErrorf(t *testing.T) {
	err := New("foo")
	err = Errorf("bar: %w", err)
	require.Equal(t, "bar: foo", err.Error())
}

//goland:noinspection GoPrintFunctions
func TestIs(t *testing.T) {
	var errFoo = New("foo")
	err := Errorf("bar: %w", errFoo)
	require.True(t, Is(err, errFoo))
	require.False(t, Is(New("baz"), errFoo))
}

func TestJoin(t *testing.T) {
	require.Equal(t, "foo\nbar", Join(New("foo"), New("bar")).Error())
}

func TestNew(t *testing.T) {
	require.Equal(t, "foo", New("foo").Error())
}

//goland:noinspection GoPrintFunctions
func TestUnwrap(t *testing.T) {
	err = Errorf("bar: %w", New("foo"))
	err = Unwrap(err)
	err = Unwrap(err)
	require.Equal(t, "foo", err.Error())
}

type errFoo struct{}

func (e errFoo) Foo() {}

func (e errFoo) Error() string { return "foo" }
