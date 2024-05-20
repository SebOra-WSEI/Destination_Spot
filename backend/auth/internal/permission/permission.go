package permission

import (
	"fmt"
	"github.com/SebastianOraczek/auth/database"
	"github.com/SebastianOraczek/auth/internal/model"
	"github.com/SebastianOraczek/auth/internal/response"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

const AdminRole string = "admin"

func User(id uint, claims jwt.MapClaims) (int, error) {
	reqUserEmail, ok := claims["Email"]
	if !ok {
		fmt.Println("Email can not be found in claims")
		return http.StatusInternalServerError, response.ErrInternalServer
	}

	var reqUser model.User
	if err := database.Db.Where("email = ?", reqUserEmail).First(&reqUser).Error; err != nil {
		fmt.Println("Requested user not found")
		return http.StatusNotFound, response.ErrUserNotFound
	}

	if reqUser.Id != id {
		fmt.Println("Password must be changed owner")
		return http.StatusForbidden, response.ErrActionNotPermitted
	}

	return 0, nil
}

func Admin(claims jwt.MapClaims) (int, error) {
	reqUserRole, ok := claims["Role"]
	if !ok {
		fmt.Println("Role can not be found in claims")
		return http.StatusInternalServerError, response.ErrInternalServer
	}

	if reqUserRole.(string) != AdminRole {
		fmt.Println("Action enabled only for admin")
		return http.StatusForbidden, response.ErrActionNotPermitted
	}

	return 0, nil
}
