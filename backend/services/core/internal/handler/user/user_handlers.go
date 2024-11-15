package user

import (
	 "fmt"
	 "github.com/SebOra-WSEI/Destination_spot/core/database"
	 "github.com/SebOra-WSEI/Destination_spot/core/internal/handler/auth"
	 "github.com/SebOra-WSEI/Destination_spot/core/internal/model"
	 userModel "github.com/SebOra-WSEI/Destination_spot/shared/model"
	 "github.com/SebOra-WSEI/Destination_spot/shared/permission"
	 "github.com/SebOra-WSEI/Destination_spot/shared/response"
	 "github.com/SebOra-WSEI/Destination_spot/shared/token"
	 "github.com/gin-gonic/gin"
	 "github.com/golang-jwt/jwt/v5"
	 "net/http"
)

func GetAll(c *gin.Context) {
	 t, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
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

func GetById(c *gin.Context) {
	 id := c.Params.ByName("id")

	 t, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
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

func Delete(c *gin.Context) {
	 id := c.Params.ByName("id")

	 t, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
	 if err != nil {
		  c.JSON(http.StatusBadRequest, response.CreateError(err))
		  return
	 }

	 if adminCode, err := permission.Admin(t.Claims.(jwt.MapClaims)); err != nil {
		  c.JSON(adminCode, response.CreateError(err))
		  return
	 }

	 var user userModel.User
	 if err := user.FindById(database.Db, id, &user); err != nil {
		  c.JSON(http.StatusNotFound, response.CreateError(err))
		  return
	 }

	 var allReservations []model.Reservation
	 if err := database.Db.Table("reservations").Select("reservations.*").
		  Where("user_id = ?", user.ID).
		  Find(&allReservations).Error; err != nil {
		  fmt.Println("Problem with fetching related reservations", err.Error())
		  c.JSON(http.StatusInternalServerError, response.Create(response.ErrInternalServer))
		  return
	 }

	 if len(allReservations) != 0 {
		  for _, r := range allReservations {
				if err := r.Delete(database.Db, &r); err != nil {
					 fmt.Println("Problem while deleting reservations during deleting spot", err.Error())
					 c.JSON(http.StatusInternalServerError, response.CreateError(response.ErrInternalServer))
					 return
				}
		  }
	 }

	 if err := database.Db.Delete(&user).Error; err != nil {
		  fmt.Println("Problem while deleting the user", err.Error())
		  c.JSON(http.StatusInternalServerError, response.Create(response.ErrProblemWhileRemovingUser))
		  return
	 }

	 res := userModel.UserResponseWithMessage{
		  Message: response.UserRemoveMsg,
		  User:    user.GetWithNoPassword(),
	 }

	 database.Db.Delete(&user)
	 c.JSON(http.StatusOK, response.Create(res))
}
