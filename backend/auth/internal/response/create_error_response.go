package response

import "github.com/gin-gonic/gin"

func CreateErrorResponse(errMsg string) gin.H {
	return gin.H{"error": errMsg}
}