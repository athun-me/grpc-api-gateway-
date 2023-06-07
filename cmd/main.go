package main

import (
	"fmt"
	"log"

	"github.com/athunlal/api-gateway/pkg/auth"
	"github.com/athunlal/api-gateway/pkg/config"
	"github.com/athunlal/api-gateway/pkg/order"
	"github.com/athunlal/api-gateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	fmt.Println("The authentication service is : ", authSvc)

	r.Run(c.Port)
}
