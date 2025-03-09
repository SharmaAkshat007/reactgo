package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.StaticFS("/", http.Dir("../web/reactgo/build"))

	router.Run(":8080")
}
