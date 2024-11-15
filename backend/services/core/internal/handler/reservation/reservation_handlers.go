package reservation

import (
	"github.com/SebOra-WSEI/Destination_spot/core/database"
	"github.com/SebOra-WSEI/Destination_spot/core/internal/handler/auth"
	"github.com/SebOra-WSEI/Destination_spot/core/internal/model"
	"github.com/SebOra-WSEI/Destination_spot/shared/permission"
	"github.com/SebOra-WSEI/Destination_spot/shared/request"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/SebOra-WSEI/Destination_spot/shared/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
)

func GetAll(c *gin.Context) {
	_, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var reservations []model.ReservationDetails
	if err := database.Db.Table("reservations").Select("reservations.*, reservations.id as reservation_id, users.*, spots.*, spots.id as spot_id").
		Joins("LEFT JOIN users ON users.id = reservations.user_id").
		Joins("LEFT JOIN spots ON spots.id = reservations.spot_id").
		Find(&reservations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	var res model.Reservation
	c.JSON(http.StatusOK, response.Create(res.GetAllWithDetails(reservations)))
}

func GetById(c *gin.Context) {
	id := c.Params.ByName("id")

	_, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var res model.Reservation
	reservation, err := res.FindByIdWithDetails(database.Db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(err))
		return
	}

	c.JSON(http.StatusOK, response.Create(reservation))
}

func Create(c *gin.Context) {
	t, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var body struct {
		UserID       uint `json:"userId"`
		SpotID       uint `json:"spotId"`
		ReservedFrom int  `json:"reservedFrom"`
		ReservedTo   int  `json:"reservedTo"`
	}

	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBodyFields(
		strconv.Itoa(body.ReservedTo), strconv.Itoa(body.ReservedFrom),
	) || body.SpotID == 0 || body.UserID == 0 {
		c.JSON(http.StatusBadRequest, response.CreateError(response.ErrEmptyFields))
		return
	}

	if _, err := permission.User(database.Db, body.UserID, t.Claims.(jwt.MapClaims)); err != nil {
		c.JSON(http.StatusBadRequest, response.CreateError(err))
		return
	}

	var allReservations []model.Reservation
	if err := database.Db.Table("reservations").Select("reservations.*").
		Where("spot_id = ?", body.SpotID).
		Where("reserved_from >= ?", body.ReservedFrom).
		Where("reserved_to <= ?", body.ReservedTo).
		Find(&allReservations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	if len(allReservations) != 0 {
		c.JSON(http.StatusBadRequest, response.CreateError(response.ErrSpotAlreadyReservedMsg))
		return
	}

	reservation := model.Reservation{
		UserID:       body.UserID,
		SpotID:       body.SpotID,
		ReservedFrom: body.ReservedFrom,
		ReservedTo:   body.ReservedTo,
	}

	if err := reservation.Create(database.Db, &reservation); err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	res := model.ReservationResponseWithMessage{
		Message:     response.ReservationCreatedMsg,
		Reservation: reservation,
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

	var reservation model.Reservation
	if err := reservation.FindById(database.Db, id, &reservation); err != nil {
		c.JSON(http.StatusNotFound, response.CreateError(err))
		return
	}

	res := model.ReservationResponseWithMessage{
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

	if err := reservation.Delete(database.Db, &reservation); err != nil {
		c.JSON(http.StatusInternalServerError, response.CreateError(err))
		return
	}

	c.JSON(http.StatusOK, response.Create(res))
}
