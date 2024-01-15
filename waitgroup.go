package main

import (
	"fmt"
	"sync"
)

func attackEvil(evil string, beeper *sync.WaitGroup) {

	fmt.Println("Attack", evil)
	beeper.Done()

}

func main() {

	var beeper sync.WaitGroup
	evilNinjas := []string{"Glenn", "Smith", "Hank", "de Mertin"}
	beeper.Add(len(evilNinjas))

	for _, evilNinja := range evilNinjas {
		go attackEvil(evilNinja, &beeper)
	}
	beeper.Wait()
	fmt.Println("Mission completed")

}
