package handler

import (
	"github.com/SebOra-WSEI/Destination_spot/core/database"
	coreModel "github.com/SebOra-WSEI/Destination_spot/core/internal/model"
	"github.com/SebOra-WSEI/Destination_spot/shared/model"
	"github.com/SebOra-WSEI/Destination_spot/shared/permission"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/SebOra-WSEI/Destination_spot/shared/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type Reservation struct {
	ID           uint `json:"id"`
	UserID       uint `json:"userId"`
	SpotID       uint `json:"spotId"`
	ReservedFrom int  `json:"reservedFrom"`
	ReservedTo   int  `json:"reservedTo"`
}

//type ReservationBody struct {
//	UserID       uint `json:"userId"`
//	SpotID       uint `json:"spotId"`
//	ReservedFrom int  `json:"reservedFrom"`
//	ReservedTo   int  `json:"reservedTo"`
//}

type ReservationResponseWithMessage struct {
	Message     string      `json:"message"`
	Reservation Reservation `json:"reservation"`
}

type ReservationWithUserAndSpot struct {
	Details Reservation          `json:"details"`
	User    model.NoPasswordUser `json:"user"`
	Spot    coreModel.Spot       `json:"spot"`
}

type ReservationResponse struct {
	Reservation ReservationWithUserAndSpot `json:"reservation"`
}

type AllReservationsResponse struct {
	Reservations []ReservationWithUserAndSpot `json:"reservations"`
}

func GetAllReservations(c *gin.Context) {
	_, err := token.Verify(c.GetHeader(AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var reservations []struct {
		ReservationId uint `json:"reservationId"`
		SpotId        uint `json:"spotId"`
		Reservation
		coreModel.Spot
		model.NoPasswordUser
	}

	if err := database.Db.Table("reservations").Select("reservations.*, reservations.id as reservation_id, users.*, spots.*, spots.id as spot_id").
		Joins("LEFT JOIN users ON users.id = reservations.user_id").
		Joins("LEFT JOIN spots ON spots.id = reservations.spot_id").
		Find(&reservations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	var res []ReservationWithUserAndSpot
	for _, r := range reservations {
		res = append(
			res, ReservationWithUserAndSpot{
				Details: Reservation{
					ID:           r.ReservationId,
					UserID:       r.UserID,
					SpotID:       r.SpotID,
					ReservedFrom: r.ReservedFrom,
					ReservedTo:   r.ReservedTo,
				},
				Spot: coreModel.Spot{
					ID:       r.SpotId,
					Location: r.Location,
				},
				User: model.NoPasswordUser{
					ID:      r.UserID,
					Email:   r.Email,
					Name:    r.Name,
					Surname: r.Surname,
					Role:    r.Role,
				},
			},
		)
	}

	var reservationsArr []ReservationWithUserAndSpot
	if len(res) > 0 {
		reservationsArr = res
	} else {
		reservationsArr = []ReservationWithUserAndSpot{}
	}

	c.JSON(http.StatusOK, AllReservationsResponse{Reservations: reservationsArr})
}

func GetReservation(c *gin.Context) {
	id := c.Params.ByName("id")

	_, err := token.Verify(c.GetHeader(AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var reservation Reservation
	if err := database.Db.First(&reservation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(response.ErrReservationNotFound))
		return
	}

	var user model.User
	if err := database.Db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(response.ErrUserNotFound))
		return
	}

	var spot coreModel.Spot
	if err := database.Db.First(&spot, id).Error; err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(response.ErrSpotNotFound))
		return
	}

	res := ReservationResponse{
		Reservation: ReservationWithUserAndSpot{
			Details: Reservation{
				ID:           reservation.ID,
				UserID:       reservation.UserID,
				SpotID:       reservation.SpotID,
				ReservedFrom: reservation.ReservedFrom,
				ReservedTo:   reservation.ReservedTo,
			},
			User: model.NoPasswordUser{
				ID:      user.ID,
				Email:   user.Email,
				Name:    user.Name,
				Surname: user.Surname,
				Role:    user.Role,
			},
			Spot: coreModel.Spot{
				ID:       spot.ID,
				Location: spot.Location,
			},
		},
	}

	c.JSON(http.StatusOK, response.Create(res))
}

func DeleteReservation(c *gin.Context) {
	id := c.Params.ByName("id")

	t, err := token.Verify(c.GetHeader(AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var reservation Reservation
	if err := database.Db.First(&reservation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(response.ErrReservationNotFound))
		return
	}

	res := ReservationResponseWithMessage{
		Message:     response.ReservationRemoveMsg,
		Reservation: reservation,
	}

	if adminCode, err := permission.Admin(t.Claims.(jwt.MapClaims)); err != nil {
		if _, err := permission.User(database.Db, reservation.UserID, t.Claims.(jwt.MapClaims)); err == nil {
			database.Db.Delete(&reservation)
			c.JSON(http.StatusOK, response.Create(res))

			return
		}

		c.JSON(adminCode, response.CreateError(err))
		return
	}

	database.Db.Delete(&reservation)
	c.JSON(http.StatusOK, response.Create(res))
}
