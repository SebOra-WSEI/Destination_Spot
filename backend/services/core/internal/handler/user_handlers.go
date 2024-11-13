package handler

import (
	"github.com/SebOra-WSEI/Destination_spot/core/database"
	"github.com/SebOra-WSEI/Destination_spot/shared/model"
	"github.com/SebOra-WSEI/Destination_spot/shared/permission"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/SebOra-WSEI/Destination_spot/shared/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	t, err := token.Verify(c.GetHeader(AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	if code, err := permission.Admin(t.Claims.(jwt.MapClaims)); err != nil {
		c.JSON(code, response.CreateError(err))
		return
	}

	var users []model.User
	if err := database.Db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	var noPasswordUsers []model.NoPasswordUser
	for _, user := range users {
		noPasswordUsers = append(
			noPasswordUsers, model.NoPasswordUser{
				ID:      user.ID,
				Email:   user.Email,
				Name:    user.Name,
				Surname: user.Surname,
				Role:    user.Role,
			},
		)
	}

	c.JSON(http.StatusOK, response.Create(model.UsersResponse{Users: noPasswordUsers}))
}
