package handler

import (
	 "fmt"
	 "github.com/SebOra-WSEI/Destination_spot/core/database"
	 "github.com/SebOra-WSEI/Destination_spot/shared/response"
	 "github.com/SebOra-WSEI/Destination_spot/shared/token"
	 "github.com/gin-gonic/gin"
	 "github.com/gin-gonic/gin/binding"
	 "net/http"
)

type Spot struct {
	 ID       uint `json:"id"`
	 Location int  `json:"location"`
}

type SpotBody struct {
	 Location int `json:"location"`
}

type SpotResponse struct {
	 Message string `json:"message"`
	 Spot    Spot   `json:"spot"`
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

func GetSpot(c *gin.Context) {
	 id := c.Params.ByName("id")

	 _, err := token.Verify(c.GetHeader(AuthorizationHeader))
	 if err != nil {
		  c.JSON(http.StatusBadRequest, response.CreateError(err))
		  return
	 }

	 var spot Spot
	 if err := database.Db.First(&spot, id).Error; err != nil {
		  c.JSON(http.StatusNotFound, response.CreateError(response.ErrSpotNotFound))
		  return
	 }

	 c.JSON(http.StatusOK, response.Create(spot))
}

func CreateSpot(c *gin.Context) {
	 var body SpotBody
	 if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		  c.JSON(http.StatusBadRequest, response.CreateError(response.ErrEmptyFields))
		  return
	 }

	 var spot Spot
	 if err := database.Db.Where("location = ?", body.Location).First(&spot).Error; err == nil {
		  c.JSON(http.StatusBadRequest, response.CreateError(response.ErrSpotAlreadyExists))
		  return
	 }

	 newSpot := Spot{
		  Location: body.Location,
	 }

	 if err := database.Db.Create(&newSpot).Error; err != nil {
		  fmt.Println("Problem creating new spot", err.Error())
		  c.JSON(http.StatusInternalServerError, response.CreateError(response.ErrProblemWhileCreatingNewSpot))
		  return
	 }

	 res := SpotResponse{
		  Message: response.SpotCreatedMsg,
		  Spot:    newSpot,
	 }

	 c.JSON(http.StatusOK, response.Create(res))
}
