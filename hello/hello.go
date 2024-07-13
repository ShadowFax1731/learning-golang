package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Rachit")
	fmt.Println(message)
}