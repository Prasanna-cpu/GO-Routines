package main

import (
	"fmt"
	"math/rand"
	"time"
)

func NinjaStar(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	numRounds := 3
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored:", score)
	}

	close(channel)
}

func main() {
	channel := make(chan string)

	go NinjaStar(channel)

	for message := range channel {
		fmt.Println(message)
	}

}
