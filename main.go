package main

import (
	"sync"

	"github.com/ItzSwirlz/pokemon-gen7/nex"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	// TODO - Add gRPC server, uncomment these
	go nex.StartAuthenticationServer()
	go nex.StartSecureServer()

	wg.Wait()
}
