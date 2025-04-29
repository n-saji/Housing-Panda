package db

import (
	"fmt"
	"housing_panda/models"

	"github.com/google/uuid"
)

func (db *DB) InsertListing(listing *models.Listings) error {
	err := db.db_conn.Create(listing).Error
	if err != nil {
		return err
	}
	return nil
}
func (db *DB) GetListingById(l_id uuid.UUID) (*models.Listings, error) {
	var listing models.Listings
	err := db.db_conn.Where("id = ?", l_id).First(&listing).Error
	if err != nil{
		fmt.Println(err)
	}
	if err != nil && err.Error() != "record not found" {
		return nil, err
	} else if err != nil && err.Error() == "record not found" {
		return nil, nil
	}
	return &listing, nil
}

func (db *DB) UpdateListingById(l_id uuid.UUID, listing *models.Listings) error {
	err := db.db_conn.Model(&models.Listings{}).Where("id = ?", l_id).Updates(listing).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) DeleteListingById(l_id uuid.UUID) error {
	err := db.db_conn.Model(&models.Listings{}).Where("id = ?", l_id).Update("is_deleted", true).Error
	if err != nil {
		return err
	}
	return nil
}
func (db *DB) GetAllListings() ([]*models.Listings, error) {
	var listings []*models.Listings
	err := db.db_conn.Where("is_deleted = ?", false).Order("created_at DESC").Find(&listings).Error
	if err != nil {
		return nil, err
	}
	return listings, nil
}

func (db *DB) GetAllListingsByUserId(u_id uuid.UUID) ([]*models.Listings, error) {
	var listings []*models.Listings
	err := db.db_conn.Where("user_id = ? AND is_deleted = ?", u_id, false).Find(&listings).Error
	if err != nil {
		return nil, err
	}
	return listings, nil
}
