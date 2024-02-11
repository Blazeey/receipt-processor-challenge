package main

import "receipt-processor-challenge/server"

func main() {
	httpServer := server.NewServer("localhost", 8080)
	go httpServer.Start()

	httpServer.WaitForShutdown()
}
