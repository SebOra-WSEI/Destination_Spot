package permission

import (
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/shared/model"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/gorm"
	"net/http"
)

const AdminRole string = "admin"

func User(db *gorm.DB, id uint, claims jwt.MapClaims) (int, error) {
	reqUserEmail, ok := claims["Email"]
	if !ok {
		fmt.Println("Email can not be found in claims")
		return http.StatusInternalServerError, response.ErrInternalServer
	}

	var reqUser model.User
	if err := reqUser.FindByEmail(db, reqUserEmail.(string), &reqUser); err != nil {
		fmt.Println("Requested user not found")
		return http.StatusNotFound, response.ErrUserNotFound
	}

	if reqUser.ID != id {
		fmt.Println("Action only permitted by owner")
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
