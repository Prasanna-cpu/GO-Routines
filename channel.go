package main

import (
	"fmt"
	"time"
)

func Attack(target string, attacked chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Throwing Ninja stars at:", target)
	attacked <- true //Send
	close(attacked)

}

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	smokeSignal := make(chan bool)
	evil := "Tommy"
	go Attack(evil, smokeSignal)
	time.Sleep(time.Second * 2)

	fmt.Println(<-smokeSignal) //Recieve

}
