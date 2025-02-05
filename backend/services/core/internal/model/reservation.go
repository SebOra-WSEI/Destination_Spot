package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/shared/model"
	userModel "github.com/SebOra-WSEI/Destination_spot/shared/model"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type Reservation struct {
	ID           uint   `json:"id"`
	UserID       uint   `json:"userId"`
	SpotID       uint   `json:"spotId"`
	ReservedFrom string `json:"reservedFrom"`
	ReservedTo   string `json:"reservedTo"`
}

type ReservationResponseWithMessage struct {
	Message     string      `json:"message"`
	Reservation Reservation `json:"reservation"`
}

type ReservationWithUserAndSpot struct {
	Details Reservation          `json:"details"`
	User    model.NoPasswordUser `json:"user"`
	Spot    Spot                 `json:"spot"`
}

type ReservationResponse struct {
	Reservation ReservationWithUserAndSpot `json:"reservation"`
}

type AllReservationsResponse struct {
	Reservations []ReservationWithUserAndSpot `json:"reservations"`
}

type ReservationDetails struct {
	ID     uint `json:"reservationId"`
	SpotId uint `json:"spotId"`
	Reservation
	Spot
	userModel.NoPasswordUser
}

type ReservationInputBody struct {
	UserID       uint   `json:"userId"`
	SpotID       uint   `json:"spotId"`
	ReservedFrom string `json:"reservedFrom"`
	ReservedTo   string `json:"reservedTo"`
}

func (r Reservation) FindById(db *gorm.DB, id string, reservation *Reservation) error {
	if err := db.First(&reservation, id).Error; err != nil {
		fmt.Println("Reservation not found:", err.Error())
		return response.ErrReservationNotFound
	}

	return nil
}

func (r Reservation) FindByIdSQL(db *sql.DB, c context.Context, id string, reservation *Reservation) error {
	rows, err := db.QueryContext(c, "SELECT * FROM reservations WHERE id = ?", id)
	defer rows.Close()

	if err != nil {
		fmt.Println("Reservation not found:", err.Error())
		return response.ErrReservationNotFound
	}

	for rows.Next() {
		if err := rows.Scan(
			&reservation.ID, &reservation.UserID, &reservation.SpotID, &reservation.ReservedFrom,
			&reservation.ReservedTo,
		); err != nil {
			fmt.Println("Problem with scanning:", err.Error())
			return response.ErrInternalServer
		}
	}

	if reservation.ID == 0 {
		return response.ErrSpotNotFound
	}

	return nil

}

func (r Reservation) FindByIdWithDetailsSQL(db *sql.DB, c context.Context, id string) (ReservationResponse, error) {
	var reservation Reservation

	rows, err := db.QueryContext(c, "SELECT * FROM reservations WHERE id = ?", id)
	defer rows.Close()

	if err != nil {
		fmt.Println("Reservations not found:", err.Error())
		return ReservationResponse{}, response.ErrReservationNotFound
	}

	for rows.Next() {
		if err := rows.Scan(
			&reservation.ID, &reservation.UserID, &reservation.SpotID, &reservation.ReservedFrom,
			&reservation.ReservedTo,
		); err != nil {
			fmt.Println("Problem with scanning:", err.Error())
			return ReservationResponse{}, response.ErrInternalServer
		}
	}

	if reservation.ID == 0 {
		fmt.Println("Reservation not found")
		return ReservationResponse{}, response.ErrReservationNotFound
	}

	var user userModel.User
	if err := user.FindByIdSQL(db, c, strconv.Itoa(int(reservation.UserID)), &user); err != nil {
		fmt.Println("User not found", err.Error())
		return ReservationResponse{}, response.ErrUserNotFound
	}

	var spot Spot
	if err := spot.FindByIdSQL(db, c, strconv.Itoa(int(reservation.SpotID)), &spot); err != nil {
		fmt.Println("Spot not found", err.Error())
		return ReservationResponse{}, response.ErrSpotNotFound
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
			User: userModel.NoPasswordUser{
				ID:      user.ID,
				Email:   user.Email,
				Name:    user.Name,
				Surname: user.Surname,
				Role:    user.Role,
			},
			Spot: Spot{
				ID:       spot.ID,
				Location: spot.Location,
			},
		},
	}

	return res, nil
}

func (r Reservation) GetAllWithDetails(reservations []ReservationDetails) []ReservationWithUserAndSpot {
	var reservationsWithDetails []ReservationWithUserAndSpot

	if len(reservations) == 0 {
		return []ReservationWithUserAndSpot{}
	}

	for _, r := range reservations {
		reservationsWithDetails = append(
			reservationsWithDetails, ReservationWithUserAndSpot{
				Details: Reservation{
					ID:           r.ID,
					UserID:       r.UserID,
					SpotID:       r.SpotID,
					ReservedFrom: r.ReservedFrom,
					ReservedTo:   r.ReservedTo,
				},
				Spot: Spot{
					ID:       r.SpotId,
					Location: r.Location,
				},
				User: userModel.NoPasswordUser{
					ID:      r.UserID,
					Email:   r.Email,
					Name:    r.Name,
					Surname: r.Surname,
					Role:    r.Role,
				},
			},
		)
	}

	return reservationsWithDetails
}

func (r Reservation) Create(db *gorm.DB, newReservation *Reservation) (int, error) {
	var spot Spot
	if err := spot.FindById(db, strconv.Itoa(int(newReservation.SpotID)), &spot); err != nil {
		fmt.Println("Selected spot not found", err.Error())
		return http.StatusBadRequest, response.ErrSpotNotFound
	}

	if status, err := handleExistingReservation(db, *newReservation); err != nil {
		fmt.Println(err)
		return status, err
	}

	if err := db.Create(&newReservation).Error; err != nil {
		fmt.Println("Problem while creating a new reservation", err.Error())
		return http.StatusInternalServerError, response.ErrProblemWhileCreatingNewReservation
	}
	return 0, nil
}

func (r Reservation) CreateSQL(db *sql.DB, c context.Context, newReservation *Reservation) (int, error) {
	var spot Spot
	if err := spot.FindByIdSQL(db, c, strconv.Itoa(int(newReservation.SpotID)), &spot); err != nil {
		fmt.Println("Selected spot not found", err.Error())
		return http.StatusBadRequest, err
	}

	if status, err := handleExistingReservationSQL(db, c, *newReservation); err != nil {
		return status, err
	}

	if _, err := db.QueryContext(
		c, "INSERT INTO reservations (user_id, spot_id, reserved_from, reserved_to) VALUE (?,?,?,?)",
		newReservation.UserID, newReservation.SpotID, newReservation.ReservedFrom, newReservation.ReservedTo,
	); err != nil {
		fmt.Println("Problem while creating a new reservation", err.Error())
		return http.StatusInternalServerError, response.ErrProblemWhileCreatingNewReservation
	}

	return 0, nil
}

func (r Reservation) Update(db *gorm.DB, newReservation *Reservation) (int, error) {
	if status, err := handleExistingReservation(db, *newReservation); err != nil {
		fmt.Println(err)
		return status, err
	}

	if err := db.Save(&newReservation).Error; err != nil {
		fmt.Println("Problem while updating a new reservation", err.Error())
		return http.StatusInternalServerError, response.ErrWhileUpdatingReservation
	}

	return 0, nil
}

func (r Reservation) UpdateSQL(db *sql.DB, c context.Context, newReservation *Reservation) (int, error) {
	if status, err := handleExistingReservationSQL(db, c, *newReservation); err != nil {
		return status, err
	}

	if _, err := db.QueryContext(
		c, "UPDATE reservations SET spot_id = ?, reserved_from = ?, reserved_to = ? WHERE id = ?",
		newReservation.SpotID, newReservation.ReservedFrom, newReservation.ReservedTo, newReservation.ID,
	); err != nil {
		fmt.Println("Problem while updating a new reservation", err.Error())
		return http.StatusInternalServerError, response.ErrWhileUpdatingReservation
	}

	return 0, nil
}

func (r Reservation) Delete(db *gorm.DB, reservation *Reservation) error {
	if err := db.Delete(&reservation).Error; err != nil {
		fmt.Println("Problem while deleting the reservation", err.Error())
		return response.ErrProblemWhileRemovingReservation
	}
	return nil
}

func handleExistingReservation(db *gorm.DB, newReservation Reservation) (int, error) {
	var allReservations []Reservation
	if err := db.Table("reservations").Select("reservations.*").
		Where("spot_id = ?", newReservation.SpotID).
		Where("reserved_from >= ?", newReservation.ReservedFrom).
		Where("reserved_to <= ?", newReservation.ReservedTo).
		Find(&allReservations).Error; err != nil {

		return http.StatusInternalServerError, err
	}

	if len(allReservations) != 0 {
		return http.StatusBadRequest, response.ErrSpotAlreadyReservedMsg
	}

	return 0, nil
}

func handleExistingReservationSQL(db *sql.DB, c context.Context, newReservation Reservation) (int, error) {
	rows, err := db.QueryContext(
		c, "SELECT * FROM reservations WHERE spot_id = ? AND reserved_from = ? AND reserved_to <= ?",
		newReservation.SpotID,
		newReservation.ReservedFrom, newReservation.ReservedTo,
	)
	defer rows.Close()

	if err != nil {
		fmt.Println("Problem while quering reservations:", err.Error())
		return http.StatusInternalServerError, response.ErrInternalServer
	}

	var existingReservations []Reservation

	for rows.Next() {
		var reservation Reservation

		if err := rows.Scan(
			&reservation.ID, &reservation.UserID, &reservation.SpotID, &reservation.ReservedFrom,
			&reservation.ReservedTo,
		); err != nil {
			fmt.Println("Problem with scanning:", err.Error())
			return http.StatusInternalServerError, response.ErrInternalServer
		}

		existingReservations = append(existingReservations, reservation)

	}

	if len(existingReservations) != 0 {
		return http.StatusBadRequest, response.ErrSpotAlreadyReservedMsg
	}

	return 0, nil
}
