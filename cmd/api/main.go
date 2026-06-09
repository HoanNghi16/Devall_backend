package main

import "github.com/HoanNghi16/Devall_backend/routes"

func main() {
	server := routes.SetupRouter()

	server.Run(":8080")
}