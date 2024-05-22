package response

import (
	"github.com/gin-gonic/gin"
)

func createResponseObject(field string, res interface{}) gin.H {
	return gin.H{field: res}
}

func CreateError(err error) gin.H {
	return createResponseObject("error", err.Error())
}

func Create(res interface{}) gin.H {
	return createResponseObject("message", res)
}
