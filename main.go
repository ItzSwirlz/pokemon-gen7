package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	// TODO - Add gRPC server, uncomment these
	// go nex.StartAuthenticationServer()
	// go nex.StartSecureServer()

	wg.Wait()
}
