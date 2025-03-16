package main

import (
	"fmt"
	"github.com/mikeyfennelly1/os--char-dev-2--u2/src/sysinfo"
)

func main() {
	port := "8989"
	numWorkers := 5
	fmt.Println("Starting server on port", port)
	sysinfo.StartServer(port, numWorkers)
}
