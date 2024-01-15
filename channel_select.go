package main

import (
	"fmt"
)

func Captain(ninja chan string, message string) {
	//time.Sleep(time.Second * 3)
	ninja <- message

}

func FairElect() {
	ninja1 := make(chan interface{})
	close(ninja1)
	ninja2 := make(chan interface{})
	close(ninja2)

	var ninja1Count, ninja2Count int

	for i := 0; i < 100; i++ {
		select {
		case <-ninja1:
			ninja1Count++
		case <-ninja2:
			ninja2Count++
		default:
			fmt.Println("Neither")
		}
	}

	fmt.Printf("ninja1Count:%d , ninja2Count:%d\n", ninja1Count, ninja2Count)
}

func main() {

	ninja1, ninja2 := make(chan string), make(chan string)

	go Captain(ninja1, "Ninja1")
	go Captain(ninja2, "Ninja2")

	select {
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
	}

	FairElect()
}
