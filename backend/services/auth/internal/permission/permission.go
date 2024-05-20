package permission

import (
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/auth/database"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/model"
	"github.com/SebOra-WSEI/Destination_spot/auth/internal/response"
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
		fmt.Println("Password must be changed by owner")
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
