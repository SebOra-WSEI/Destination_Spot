package handler

import (
	"github.com/SebOra-WSEI/Destination_spot/core/database"
	userModel "github.com/SebOra-WSEI/Destination_spot/shared/model"
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

	var users []userModel.User
	if err := database.Db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	var noPasswordUsers []userModel.NoPasswordUser
	for _, user := range users {
		noPasswordUsers = append(
			noPasswordUsers, user.GetWithNoPassword(),
		)
	}

	c.JSON(http.StatusOK, response.Create(userModel.AllUsersResponse{Users: noPasswordUsers}))
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")

	t, err := token.Verify(c.GetHeader(AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var user userModel.User
	if err := user.FindById(database.Db, id, &user); err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(response.ErrUserNotFound))
		return
	}

	noPasswordUser := userModel.UserResponse{
		User: user.GetWithNoPassword(),
	}

	if adminCode, err := permission.Admin(t.Claims.(jwt.MapClaims)); err != nil {
		if _, err := permission.User(database.Db, user.ID, t.Claims.(jwt.MapClaims)); err == nil {
			c.JSON(http.StatusOK, response.Create(noPasswordUser))
			return
		}

		c.JSON(adminCode, response.CreateError(err))
		return
	}

	c.JSON(http.StatusOK, response.Create(noPasswordUser))
}
