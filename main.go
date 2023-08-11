package main

import (
	"github.com/YottaHub/ticktick/service"
  "github.com/YottaHub/ticktick/database"
	"github.com/gin-gonic/gin"
  "fmt"
  "os"
)

func main() {
	go service.RunMessageServer()

  if err := databse.Init(); err != nil {
    fmt.Println("Failed to connect db")
    os.Exit(-1)
  }

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
