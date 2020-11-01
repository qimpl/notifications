package main

import (
	_ "github.com/qimpl/notifications/docs"
	"github.com/qimpl/notifications/router"

	_ "github.com/joho/godotenv/autoload"
)

// @title Notifications API
// @version 0.1.0
// @BasePath /api/v1
func main() {
	router.CreateRouter()
}
