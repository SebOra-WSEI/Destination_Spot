package permission

import (
	"fmt"
	"github.com/SebastianOraczek/auth/database"
	"github.com/SebastianOraczek/auth/internal/model"
	"github.com/SebastianOraczek/auth/internal/response"
	"github.com/golang-jwt/jwt/v5"
)

func Get(id uint, claims jwt.MapClaims) error {
	reqUserEmail, ok := claims["Email"]
	if !ok {
		fmt.Println("Role can not be found in claims")
		return fmt.Errorf(response.InternalServerErrMsg)
	}

	var reqUser model.User
	if err := database.Db.Where("email = ?", reqUserEmail).First(&reqUser).Error; err != nil {
		fmt.Println("Requested user not found")
		return fmt.Errorf(response.UserNotFoundErrMsg)
	}

	if reqUser.Id != id {
		fmt.Println("Password must be changed by owner or admin")
		return fmt.Errorf(response.ActionNotPermittedErrMsg)
	}

	return nil
}
