package main

import (
	"Templates/server"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("No Argments needed (:")
		return
	}
	server.ServerStart()
}