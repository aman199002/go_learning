package main

import "fmt"

func main(){

	messages := make(chan string,3)

	messages <- "message1"
	messages <- "message2"
	messages <- "message3"

	msg1 := <-messages
	msg2 := <-messages

	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(<-messages)
}