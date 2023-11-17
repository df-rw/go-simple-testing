package message

import "fmt"

// Message is a container for two variables, one that is not exported and one
// that is exported.
type Message struct {
	unexportedMessage string
	ExportedMessage   string
}

// ExportedFunction is an exported function.
func ExportedFunction(m *Message) string {
	return fmt.Sprintf("ExportedFunction: %q %q", m.unexportedMessage, m.ExportedMessage)
}

// unexportedFunction is an unexported function.
func unexportedFunction(m *Message) string {
	return fmt.Sprintf("unexportedFunction: %q %q", m.unexportedMessage, m.ExportedMessage)
}

// ExportedMethod is an exported method.
func (m *Message) ExportedMethod() string {
	return fmt.Sprintf("ExportedMethod: %q %q", m.unexportedMessage, m.ExportedMessage)
}

// unexportedMethod is an unexpoted method.
func (m *Message) unexportedMethod() string {
	return fmt.Sprintf("unexportedMethod: %q %q", m.unexportedMessage, m.ExportedMessage)
}

// New returns a new Message.
func New(unexported string, exported string) *Message {
	return &Message{
		unexportedMessage: unexported,
		ExportedMessage:   exported,
	}
}
