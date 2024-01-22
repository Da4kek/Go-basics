package main

import "fmt"

func hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func main() {
	fmt.Println(hello("Dark"))
}
