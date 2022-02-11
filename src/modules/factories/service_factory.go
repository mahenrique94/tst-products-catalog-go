package factories

import (
	"github.com/gin-gonic/gin"
)

func ServiceFactory() *gin.Engine {
	r := gin.Default()
	return r
}
