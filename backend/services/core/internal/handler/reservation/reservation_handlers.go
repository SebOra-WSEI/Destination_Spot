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
)

func GetAll(c *gin.Context) {
	 _, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
	 if err != nil {
		  c.JSON(http.StatusBadRequest, response.CreateError(err))
		  return
	 }

	 var reservation model.Reservation
	 var reservations []model.ReservationDetails

	 if err := reservation.GetAllSQL(database.DbSQL, c, &reservations); err != nil {
		  c.JSON(http.StatusInternalServerError, response.CreateError(err))
		  return
	 }

	 c.JSON(
		  http.StatusOK,
		  response.Create(
				struct {
					 Reservations []model.ReservationWithUserAndSpot `json:"reservations"`
				}{
					 Reservations: reservation.GetAllWithDetails(reservations),
				},
		  ),
	 )
}

func GetById(c *gin.Context) {
	 id := c.Params.ByName("id")

	 _, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
	 if err != nil {
		  c.JSON(http.StatusBadRequest, response.CreateError(err))
		  return
	 }

	 var res model.Reservation

	 reservation, err := res.FindByIdWithDetailsSQL(database.DbSQL, c, id)
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

	 var body model.ReservationInputBody

	 if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBodyFields(
		  body.ReservedTo, body.ReservedFrom,
	 ) || body.SpotID == 0 || body.UserID == 0 {
		  c.JSON(http.StatusBadRequest, response.CreateError(response.ErrEmptyFields))
		  return
	 }

	 if _, err := permission.UserSQL(database.DbSQL, c, body.UserID, t.Claims.(jwt.MapClaims)); err != nil {
		  c.JSON(http.StatusBadRequest, response.CreateError(err))
		  return
	 }

	 reservation := model.Reservation{
		  UserID:       body.UserID,
		  SpotID:       body.SpotID,
		  ReservedFrom: body.ReservedFrom,
		  ReservedTo:   body.ReservedTo,
	 }

	 if status, err := reservation.CreateSQL(database.DbSQL, c, &reservation); err != nil {
		  c.JSON(status, response.CreateError(err))
		  return
	 }

	 res := model.ReservationResponseWithMessage{
		  Message:     response.ReservationCreatedMsg,
		  Reservation: reservation,
	 }

	 c.JSON(http.StatusCreated, response.Create(res))
}

func Update(c *gin.Context) {
	 id := c.Params.ByName("id")

	 t, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
	 if err != nil {
		  c.JSON(http.StatusBadRequest, response.CreateError(err))
		  return
	 }

	 var body model.ReservationInputBody
	 if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil || request.HandleEmptyBodyFields(
		  body.ReservedTo, body.ReservedFrom,
	 ) || body.SpotID == 0 || body.UserID == 0 {
		  c.JSON(http.StatusBadRequest, response.CreateError(response.ErrEmptyFields))
		  return
	 }

	 var reservation model.Reservation
	 if err := reservation.FindByIdSQL(database.DbSQL, c, id, &reservation); err != nil {
		  c.JSON(http.StatusNotFound, response.CreateError(err))
		  return
	 }

	 if status, err := permission.UserSQL(database.DbSQL, c, body.UserID, t.Claims.(jwt.MapClaims)); err != nil {
		  c.JSON(status, response.CreateError(err))
		  return
	 }

	 updatedReservation := model.Reservation{
		  ID:           reservation.ID,
		  UserID:       reservation.UserID,
		  SpotID:       body.SpotID,
		  ReservedFrom: body.ReservedFrom,
		  ReservedTo:   body.ReservedTo,
	 }

	 if status, err := updatedReservation.UpdateSQL(database.DbSQL, c, &updatedReservation); err != nil {
		  c.JSON(status, response.CreateError(err))
		  return
	 }

	 res := model.ReservationResponseWithMessage{
		  Message:     response.ReservationUpdatedMsg,
		  Reservation: updatedReservation,
	 }

	 c.JSON(http.StatusOK, response.Create(res))
}

func Delete(c *gin.Context) {
	 id := c.Params.ByName("id")

	 t, err := token.Verify(c.GetHeader(auth.AuthorizationHeader))
	 if err != nil {
		  c.JSON(http.StatusBadRequest, response.CreateError(err))
		  return
	 }

	 var reservation model.Reservation
	 if err := reservation.FindByIdSQL(database.DbSQL, c, id, &reservation); err != nil {
		  c.JSON(http.StatusNotFound, response.CreateError(err))
		  return
	 }

	 res := model.ReservationResponseWithMessage{
		  Message:     response.ReservationRemoveMsg,
		  Reservation: reservation,
	 }

	 if status, err := permission.Admin(t.Claims.(jwt.MapClaims)); err != nil {
		  if _, err := permission.UserSQL(database.DbSQL, c, reservation.UserID, t.Claims.(jwt.MapClaims)); err == nil {
				if err := reservation.DeleteSQL(database.DbSQL, c, &reservation); err != nil {
					 c.JSON(status, response.CreateError(err))
				}

				c.JSON(http.StatusOK, response.Create(res))
				return
		  }

		  c.JSON(status, response.CreateError(err))
		  return
	 }

	 if err := reservation.DeleteSQL(database.DbSQL, c, &reservation); err != nil {
		  c.JSON(http.StatusInternalServerError, response.CreateError(err))
		  return
	 }

	 c.JSON(http.StatusOK, response.Create(res))
}
