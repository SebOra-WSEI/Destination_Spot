package handler

import (
	"github.com/SebOra-WSEI/Destination_spot/core/database"
	"github.com/SebOra-WSEI/Destination_spot/core/internal/model"
	"github.com/SebOra-WSEI/Destination_spot/shared/permission"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/SebOra-WSEI/Destination_spot/shared/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func GetAllSpots(c *gin.Context) {
	_, err := token.Verify(c.GetHeader(AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var spots []model.Spot
	if err := database.Db.Find(&spots).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	c.JSON(http.StatusOK, response.Create(model.AllSpotsResponse{Spots: spots}))
}

func GetSpot(c *gin.Context) {
	id := c.Params.ByName("id")

	_, err := token.Verify(c.GetHeader(AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var spot model.Spot
	if err := spot.FindById(database.Db, id, &spot); err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(err))
		return
	}

	c.JSON(http.StatusOK, response.Create(model.SpotResponse{Spot: spot}))
}

func CreateSpot(c *gin.Context) {
	t, err := token.Verify(c.GetHeader(AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	if code, err := permission.Admin(t.Claims.(jwt.MapClaims)); err != nil {
		c.JSON(code, response.CreateError(err))
		return
	}

	var body model.SpotBody
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(response.ErrEmptyFields))
		return
	}

	var spot model.Spot
	if err := spot.FindByLocation(database.Db, body.Location, &spot); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	newSpot := model.Spot{
		Location: body.Location,
	}

	if err := newSpot.Create(database.Db, &newSpot); err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	res := model.SpotResponseWithMessage{
		Message: response.SpotCreatedMsg,
		Spot:    newSpot,
	}

	c.JSON(http.StatusOK, response.Create(res))
}

func DeleteSpot(c *gin.Context) {
	id := c.Params.ByName("id")

	t, err := token.Verify(c.GetHeader(AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	if code, err := permission.Admin(t.Claims.(jwt.MapClaims)); err != nil {
		c.JSON(code, response.CreateError(err))
		return
	}

	var spot model.Spot
	if err := spot.FindById(database.Db, id, &spot); err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(err))
		return
	}

	if err := spot.Delete(database.Db, &spot); err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	res := model.SpotResponseWithMessage{
		Message: response.SpotRemoveMsg,
		Spot:    spot,
	}

	c.JSON(http.StatusOK, response.Create(res))
}
