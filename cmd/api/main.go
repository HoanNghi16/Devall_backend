package main

import (
	"github.com/HoanNghi16/Devall_backend/internal/database"
	"github.com/HoanNghi16/Devall_backend/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	server := routes.SetupRouter(db)
	server.Run(":8080")
}