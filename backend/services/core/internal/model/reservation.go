package model

import (
	 "fmt"
	 "github.com/SebOra-WSEI/Destination_spot/shared/model"
	 userModel "github.com/SebOra-WSEI/Destination_spot/shared/model"
	 "github.com/SebOra-WSEI/Destination_spot/shared/response"
	 "github.com/jinzhu/gorm"
	 "strconv"
)

type Reservation struct {
	 ID           uint `json:"id"`
	 UserID       uint `json:"userId"`
	 SpotID       uint `json:"spotId"`
	 ReservedFrom int  `json:"reservedFrom"`
	 ReservedTo   int  `json:"reservedTo"`
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

func (r Reservation) FindById(db *gorm.DB, id string, reservation *Reservation) error {
	 if err := db.First(&reservation, id).Error; err != nil {
		  fmt.Println("Reservation not found:", err.Error())
		  return response.ErrReservationNotFound
	 }

	 return nil
}

func (r Reservation) FindByIdWithDetails(db *gorm.DB, id string) (ReservationResponse, error) {
	 var reservation Reservation
	 if err := db.First(&reservation, id).Error; err != nil {
		  fmt.Println("Reservation not found", err.Error())
		  return ReservationResponse{}, response.ErrReservationNotFound
	 }

	 var user userModel.User
	 if err := db.First(&user, reservation.UserID).Error; err != nil {
		  fmt.Println("User not found", err.Error())
		  return ReservationResponse{}, response.ErrUserNotFound
	 }

	 var spot Spot
	 if err := db.First(&spot, reservation.SpotID).Error; err != nil {
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

func (r Reservation) Create(db *gorm.DB, newReservation *Reservation) error {
	 var spot Spot
	 if err := spot.FindById(db, strconv.Itoa(int(newReservation.SpotID)), &spot); err != nil {
		  fmt.Println("Selected spot not found", err.Error())
		  return response.ErrSpotNotFound
	 }

	 if err := db.Create(&newReservation).Error; err != nil {
		  fmt.Println("Problem while creating a new reservation", err.Error())
		  return response.ErrProblemWhileCreatingNewReservation
	 }
	 return nil
}

func (r Reservation) Delete(db *gorm.DB, reservation *Reservation) error {
	 if err := db.Delete(&reservation).Error; err != nil {
		  fmt.Println("Problem while deleting the reservation", err.Error())
		  return response.ErrProblemWhileRemovingReservation
	 }
	 return nil
}
