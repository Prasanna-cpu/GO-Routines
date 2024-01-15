package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock  sync.Mutex
	count int
)

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}

func util() {
	itr := 1000

	for i := 0; i < itr; i++ {
		go increment()
	}
	time.Sleep(10 * time.Second)
	fmt.Println("Resulting count is :", count)
}

func main() {

	util()

}
