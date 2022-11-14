package main

import "fmt"

func changeValue(str *MyMessage) {
	fmt.Println(str.message)
	fmt.Println((*str).message)

	(*str).message = "changed!"

	fmt.Println(str.message)
	fmt.Println((*str).message)
}

type MyMessage struct {
	message string
}

func main() {
	x := MyMessage{
		message: "hello",
	}

	changeValue(&x)
}
