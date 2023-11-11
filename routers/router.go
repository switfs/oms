package routers

import (
	"github.com/gin-gonic/gin"
)

type GroupRouter struct {
}

func InitRouter() *gin.Engine {
	routers := gin.New()

	//routers.Use(gin.Logger())

	return routers

}
