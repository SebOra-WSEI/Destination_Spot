package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/shared/response"
	"github.com/jinzhu/gorm"
)

type Spot struct {
	ID       uint `json:"id"`
	Location int  `json:"location"`
}

type SpotResponseWithMessage struct {
	Message string `json:"message"`
	Spot    Spot   `json:"spot"`
}

type SpotResponse struct {
	Spot Spot `json:"spot"`
}

type AllSpotsResponse struct {
	Spots []Spot `json:"spots"`
}

func (u Spot) FindById(db *gorm.DB, id string, spot *Spot) error {
	if err := db.First(&spot, id).Error; err != nil {
		fmt.Println("Spot not found:", err.Error())
		return response.ErrSpotNotFound
	}

	return nil
}

func (u Spot) FindByIdSQL(db *sql.DB, c context.Context, id string, spot *Spot) error {
	rows, err := db.QueryContext(c, "SELECT * FROM spots WHERE id = ?", id)
	defer rows.Close()

	if err != nil {
		fmt.Println("Spot not found:", err.Error())
		return response.ErrSpotNotFound
	}

	for rows.Next() {
		if err := rows.Scan(&spot.ID, &spot.Location); err != nil {
			fmt.Println("Problem with scanning:", err.Error())
			return response.ErrInternalServer
		}
	}

	if spot.ID == 0 {
		return response.ErrSpotNotFound
	}

	return nil
}

func (u Spot) FindAllSQL(db *sql.DB, c context.Context, spots *[]Spot) error {
	rows, err := db.QueryContext(c, "SELECT * FROM spots")
	defer rows.Close()

	if err != nil {
		fmt.Println("Spots not found:", err.Error())
		return response.ErrUserNotFound
	}

	for rows.Next() {
		var spot Spot

		if err := rows.Scan(&spot.ID, &spot.Location); err != nil {
			fmt.Println("Problem with scanning:", err.Error())
			return response.ErrInternalServer
		}

		*spots = append(*spots, spot)
	}

	if len(*spots) == 0 {
		return response.ErrUsersNotFound
	}

	return nil
}

func (u Spot) FindByLocationSQL(db *sql.DB, c context.Context, location int, spot *Spot) error {
	rows, err := db.QueryContext(c, "SELECT * FROM spots WHERE location = ?", location)
	defer rows.Close()

	if err != nil {
		fmt.Println("Spot not found:", err.Error())
		return response.ErrSpotNotFound
	}

	for rows.Next() {
		if err := rows.Scan(&spot.ID, &spot.Location); err != nil {
			fmt.Println("Problem with scanning:", err.Error())
			return response.ErrInternalServer
		}
	}

	if spot.ID != 0 {
		return response.ErrSpotAlreadyExists
	}

	return nil
}

func (u Spot) CreateSQL(db *sql.DB, c context.Context, newSpot *Spot) error {
	rows, err := db.QueryContext(c, "INSERT INTO spots (location) VALUE (?)", newSpot.Location)
	defer rows.Close()

	if err != nil {
		fmt.Println("Problem while creating a new spot", err.Error())
		return response.ErrProblemWhileCreatingNewSpot
	}

	for rows.Next() {
		if err := rows.Scan(&newSpot.ID, &newSpot.Location); err != nil {
			fmt.Println("Problem with scanning:", err.Error())
			return response.ErrInternalServer
		}
	}

	return nil
}

func (u Spot) DeleteSQL(db *sql.DB, c context.Context, spot *Spot) error {
	rows, err := db.QueryContext(c, "DELETE FROM spots WHERE id = ?", spot.ID)
	defer rows.Close()

	if err != nil {
		fmt.Println("Problem while deleting the spot", err.Error())
		return response.ErrProblemWhileRemovingSpot
	}

	for rows.Next() {
		if err := rows.Scan(&spot.ID, &spot.Location); err != nil {
			fmt.Println("Problem with scanning:", err.Error())
			return response.ErrInternalServer
		}
	}

	return nil
}
