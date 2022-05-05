package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3) // deadlock
	go randSleep(wg, "first:", 4, 3)
	go randSleep(wg, "second:", 4, 3)
	wg.Wait()
}

func randSleep(wg *sync.WaitGroup, name string, limit int, sleep int) {
	defer wg.Done()
	for i := 1; i <= limit; i++ {
		fmt.Println(name, rand.Intn(i))
		time.Sleep(time.Duration(sleep * int(time.Second)))

	}
}
