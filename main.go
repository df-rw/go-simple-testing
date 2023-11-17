package main

import (
	"fmt"
	"test/pkg/message"
)

func main() {
	m := message.New("hello", "world")

	fmt.Println(message.ExportedFunction(m))
	fmt.Println(m.ExportedMethod())
	fmt.Println(m.ExportedMessage)
}
