package model

import (
	 "fmt"
	 "github.com/SebOra-WSEI/Destination_spot/shared/response"
	 "github.com/jinzhu/gorm"
)

type Spot struct {
	 ID       uint `json:"id"`
	 Location int  `json:"location"`
}

type SpotBody struct {
	 Location int `json:"location"`
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

func (u Spot) FindByLocation(db *gorm.DB, location int, spot *Spot) error {
	 if err := db.Where("location = ?", location).First(&spot).Error; err == nil {
		  return response.ErrSpotAlreadyExists
	 }

	 return nil
}

func (u Spot) Create(db *gorm.DB, newSpot *Spot) error {
	 if err := db.Create(&newSpot).Error; err != nil {
		  fmt.Println("Problem while creating a new spot", err.Error())
		  return response.ErrProblemWhileCreatingNewSpot
	 }
	 return nil
}

func (u Spot) Delete(db *gorm.DB, spot *Spot) error {
	 if err := db.Delete(&spot).Error; err != nil {
		  fmt.Println("Problem while deleting the spot", err.Error())
		  return response.ErrProblemWhileRemovingNewSpot
	 }
	 return nil
}
