package routes

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/athunlal/api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

func FindOne(ctx *gin.Context, c pb.ProductServiceClient) {
	
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid ID: %s", idStr))
		return
	}

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusFound, &res)
}
