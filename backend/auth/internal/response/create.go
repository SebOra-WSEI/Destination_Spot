package response

import (
	 "github.com/gin-gonic/gin"
)

func createResponseObject(field string, res interface{}) gin.H {
	 return gin.H{field: res}
}

func CreateError(errMsg string) gin.H {
	 return createResponseObject("error", errMsg)
}

func Create(res interface{}) gin.H {
	 return createResponseObject("response", res)
}
