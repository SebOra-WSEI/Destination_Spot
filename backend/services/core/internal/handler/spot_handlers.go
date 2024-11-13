package handler

import (
	 "github.com/SebOra-WSEI/Destination_spot/core/database"
	 "github.com/SebOra-WSEI/Destination_spot/shared/response"
	 "github.com/SebOra-WSEI/Destination_spot/shared/token"
	 "github.com/gin-gonic/gin"
	 "net/http"
)

type Spot struct {
	 ID       uint `json:"id"`
	 Location int  `json:"location"`
}

type AllSpotsResponse struct {
	 Spots []Spot `json:"spots"`
}

func GetAllSpots(c *gin.Context) {
	 _, err := token.Verify(c.GetHeader(AuthorizationHeader))
	 if err != nil {
		  c.JSON(http.StatusBadRequest, response.CreateError(err))
		  return
	 }

	 var spots []Spot
	 if err := database.Db.Find(&spots).Error; err != nil {
		  c.JSON(http.StatusInternalServerError, response.CreateError(err))
		  return
	 }

	 c.JSON(http.StatusOK, response.Create(AllSpotsResponse{Spots: spots}))
}
