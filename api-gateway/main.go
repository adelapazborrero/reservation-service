package main

import (
	"os"

	"github.com/Abeldlp/reservation-service/api-gateway/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("PROD_ENV")

	if env == "" {
		godotenv.Load(".env")
	}
	r := gin.Default()

	r.Any("/*proxyPath", handlers.ProxyRequestToServer)

	r.Run(":8080")
}
