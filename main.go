package main

import (
	"housing_panda/config"
	"housing_panda/handlers"
)

func main() {
	config.Init()
	db_conn := config.ConnectDB()
	if db_conn == nil {
		panic("Failed to connect to database")
	}
	config.RunGooseMigration(db_conn)

	handlers := handlers.NewHandlers(db_conn)

	route := handlers.GetRouter()
	err := route.Run(config.PORT)
	if err != nil {
		panic("Failed to start server")
	}

	defer config.CloseDB(db_conn)
}
