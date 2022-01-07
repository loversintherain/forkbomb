package main

import (
	"flag"
	"sync"
)


func bomb(wg *sync.WaitGroup){
	wg.Add(1)
	go bomb(wg)
	for {

	}
}

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go bomb(&wg)
	wg.Wait()
}

