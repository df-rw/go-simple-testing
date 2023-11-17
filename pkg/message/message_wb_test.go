// Package message_wb_test performs white-box testing of package message.
package message

import (
	"testing"
)

func TestUnexportedFunction(t *testing.T) {
	m := New("hello", "world")

	got := unexportedFunction(m)
	want := "unexportedFunction: \"hello\" \"world\""

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestUnexportedMethod(t *testing.T) {
	m := New("hello", "world")

	got := m.unexportedMethod()
	want := "unexportedMethod: \"hello\" \"world\""

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestUnexportedMessage(t *testing.T) {
	m := New("hello", "world")

	got := m.unexportedMessage
	want := "hello"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
