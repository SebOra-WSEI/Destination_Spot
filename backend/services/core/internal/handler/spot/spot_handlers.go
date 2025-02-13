package spot

import (
	"github.com/SebOra-WSEI/Destination_spot/core/database"
	"github.com/SebOra-WSEI/Destination_spot/core/internal/handler/auth"
	"github.com/SebOra-WSEI/Destination_spot/core/internal/model"
	"github.com/SebOra-WSEI/Destination_spot/shared/permission"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/SebOra-WSEI/Destination_spot/shared/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func GetAll(c *gin.Context) {
	_, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var spot model.Spot
	var spots []model.Spot
	if err := spot.FindAllSQL(database.DbSQL, c, &spots); err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	c.JSON(http.StatusOK, response.Create(model.AllSpotsResponse{Spots: spots}))
}

func GetById(c *gin.Context) {
	id := c.Params.ByName("id")

	_, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var spot model.Spot
	if err := spot.FindByIdSQL(database.DbSQL, c, id, &spot); err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(err))
		return
	}

	c.JSON(http.StatusOK, response.Create(model.SpotResponse{Spot: spot}))
}

func Create(c *gin.Context) {
	t, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	if status, err := permission.Admin(t.Claims.(jwt.MapClaims)); err != nil {
		c.JSON(status, response.CreateError(err))
		return
	}

	var body struct {
		Location int `json:"location"`
	}

	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(response.ErrEmptyFields))
		return
	}

	var spot model.Spot
	if err := spot.FindByLocationSQL(database.DbSQL, c, body.Location, &spot); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	newSpot := model.Spot{
		Location: body.Location,
	}

	if err := newSpot.CreateSQL(database.DbSQL, c, &newSpot); err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	res := model.SpotResponseWithMessage{
		Message: response.SpotCreatedMsg,
		Spot:    newSpot,
	}

	c.JSON(http.StatusCreated, response.Create(res))
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	t, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	if status, err := permission.Admin(t.Claims.(jwt.MapClaims)); err != nil {
		c.JSON(status, response.CreateError(err))
		return
	}

	var spot model.Spot
	if err := spot.FindByIdSQL(database.DbSQL, c, id, &spot); err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(err))
		return
	}

	if err := spot.DeleteSQL(database.DbSQL, c, &spot); err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	res := model.SpotResponseWithMessage{
		Message: response.SpotRemoveMsg,
		Spot:    spot,
	}

	c.JSON(http.StatusOK, response.Create(res))
}
