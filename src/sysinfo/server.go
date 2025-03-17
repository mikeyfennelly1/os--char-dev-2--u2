// sysinfo server that provides sysinfo over HTTP to clients

package sysinfo

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

// Simple HTTP handler for the sysinfo response
func sysInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with some sysinfo or custom data
	w.Write([]byte("Sysinfo: Everything is running smoothly!\n"))
}

// Create a custom listener that serves HTTP requests over a single connection
type singleConnListener struct {
	conn net.Conn
}

func (l *singleConnListener) Accept() (net.Conn, error) {
	return l.conn, nil // Return the single connection
}

func (l *singleConnListener) Close() error {
	return l.conn.Close() // Close the connection when done
}

func (l *singleConnListener) Addr() net.Addr {
	return l.conn.LocalAddr() // Return the local address of the connection
}

func worker(taskChan chan net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for conn := range taskChan {
		handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Create a custom listener using the net.Conn
	listener := &singleConnListener{conn: conn}

	// Create a HTTP server with the sysInfoHandler
	server := &http.Server{
		Handler: http.HandlerFunc(sysInfoHandler), // Use the sysInfoHandler for requests
	}

	// Serve HTTP requests via the custom listener (over the single connection)
	err := server.Serve(listener)
	if err != nil {
		// Log error if something goes wrong while serving
		fmt.Println("Error serving HTTP:", err)
	}
}

func createWorkerPool(numWorkers int) chan net.Conn {
	// Create a channel for goroutines to communicate
	taskChan := make(chan net.Conn)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		// Add a wait to the wait group
		wg.Add(1)
		// Spin off new goroutine (thread), adding it to the wait group
		go worker(taskChan, &wg)
	}

	// New goroutine to provide wait context for workers
	go func() {
		wg.Wait()
		// Close channel when all threads have completed
		close(taskChan)
	}()
	return taskChan
}

// StartServer listens on a given port and starts the worker pool
func StartServer(port string, numWorkers int) {
	// Listen for incoming connections on the specified port
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	// Create the worker pool
	taskChan := createWorkerPool(numWorkers)

	// Accept incoming connections and send them to the worker pool
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Send the accepted connection to the task channel
		taskChan <- conn
	}
}
