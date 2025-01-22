package main

import (
	"example/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
