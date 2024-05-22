package permission

import (
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/shared/message"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/gorm"
	"net/http"
)

type UserModel struct {
	Id       uint
	Email    string
	Password string
	Role     string
	Name     string
	Surname  string
}

const AdminRole string = "admin"

func User(db *gorm.DB, id uint, claims jwt.MapClaims) (int, error) {
	reqUserEmail, ok := claims["Email"]
	if !ok {
		fmt.Println("Email can not be found in claims")
		return http.StatusInternalServerError, message.ErrInternalServer
	}

	var reqUser UserModel
	if err := db.Where("email = ?", reqUserEmail).First(&reqUser).Error; err != nil {
		fmt.Println("Requested user not found")
		return http.StatusNotFound, message.ErrUserNotFound
	}

	if reqUser.Id != id {
		fmt.Println("Password must be changed by owner")
		return http.StatusForbidden, message.ErrActionNotPermitted
	}

	return 0, nil
}

func Admin(claims jwt.MapClaims) (int, error) {
	reqUserRole, ok := claims["Role"]
	if !ok {
		fmt.Println("Role can not be found in claims")
		return http.StatusInternalServerError, message.ErrInternalServer
	}

	if reqUserRole.(string) != AdminRole {
		fmt.Println("Action enabled only for admin")
		return http.StatusForbidden, message.ErrActionNotPermitted
	}

	return 0, nil
}
