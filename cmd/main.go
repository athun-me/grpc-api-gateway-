package main

import (
	"fmt"
	"log"

	"github.com/athunlal/api-gateway/pkg/auth"
	"github.com/athunlal/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	fmt.Println("The authentication serivce is : ", authSvc)

	r.Run(c.Port)
}
