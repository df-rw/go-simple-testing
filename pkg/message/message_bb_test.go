// Package message_bb_test performs black-box testing of package message.
package message_test

import (
	"test/pkg/message"
	"testing"
)

func TestExportedFunction(t *testing.T) {
	m := message.New("hello", "world")

	got := message.ExportedFunction(m)
	want := "ExportedFunction: \"hello\" \"world\""

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestExportedMethod(t *testing.T) {
	m := message.New("hello", "world")

	got := m.ExportedMethod()
	want := "ExportedMethod: \"hello\" \"world\""

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestExportedMessage(t *testing.T) {
	m := message.New("hello", "world")

	got := m.ExportedMessage
	want := "world"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
