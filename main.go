package main

import (
	"bai2/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.Registeroutes(r)
	fmt.Println("Chạy trên cổng 9999")
	r.Run(":9999")
}
