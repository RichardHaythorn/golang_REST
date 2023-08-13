package main

import (
	"sync"

	"github.com/RichardHaythorn/golang_REST/api"
	"github.com/RichardHaythorn/golang_REST/database"
)

func main() {
	var wg sync.WaitGroup
	go api.Main()
	wg.Add(1)
	go database.Main()
	wg.Add(1)

	wg.Wait()

}
