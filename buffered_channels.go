package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "Message"
	channel <- "Second "
	fmt.Println(<-channel)
	fmt.Println(<-channel)

}
