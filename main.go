package main

import (
	_ "github.com/qimpl/APP_NAME/docs"
	"github.com/qimpl/APP_NAME/router"

	_ "github.com/joho/godotenv/autoload"
)

// @title APP_NAME API
// @version 0.1.0
// @BasePath /api/v1
func main() {
	router.CreateRouter()
}
