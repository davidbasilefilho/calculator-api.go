package main

import "log"

func main() {
	server := NewAPIServer(":8080")
	err := server.Run()

	if err != nil {
		log.Println("Error: ", err)
	}
}
