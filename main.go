package main

import (
	"sync"

	"github.com/RichardHaythorn/golang_REST/api"
)

func main() {
	var wg sync.WaitGroup
	go api.Main()
	wg.Add(1)

	wg.Wait()

}
