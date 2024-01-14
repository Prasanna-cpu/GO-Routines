package main

import (
	"fmt"
	"math/rand"
	"time"
)

func NinjaStar(channel chan string, numRounds int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored:", score)
	}

	close(channel)
}

func main() {
	channel := make(chan string)
	numRounds := 3
	go NinjaStar(channel, numRounds)
	for i := 0; i < numRounds; i++ {
		fmt.Println(<-channel)
	}
	fmt.Println(<-channel)
}
