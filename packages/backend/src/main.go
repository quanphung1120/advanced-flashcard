package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  router := gin.Default()

  router.GET("/ping", func(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H {
      "message": "pong",
    })
  })

  router.Run()
}