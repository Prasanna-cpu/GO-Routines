package main

import (
	"fmt"
	"time"
)

func attack(target string) {
	fmt.Println("Throwing Ninja stars at:", target)
	time.Sleep(time.Second)

}

func main() {

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	evilNinjas := []string{"Tom", "John", "Kuhn", "Rasputin", "Joe"}

	for _, v := range evilNinjas {
		go attack(v)
	}
	time.Sleep(2 * time.Second)

}
