package main

import (
	"runtime"
	"sync"
)


func bomb(wg *sync.WaitGroup) {
	wg.Add(1)
	go bomb(wg)
	for {

	}
	wg.Done()
}

func main() {
	n := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go bomb(&wg)
	}
	wg.Wait()
}
