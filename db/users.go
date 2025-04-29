package db

import (
	"housing_panda/models"

	"github.com/google/uuid"
)

func (db *DB) InsertUser(user *models.Users) error {
	err := db.db_conn.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
func (db *DB) GetUserById(u_id uuid.UUID) (*models.Users, error) {
	var user models.Users
	err := db.db_conn.Where("id = ?", u_id).First(&user).Error
	if err != nil && err.Error() != "record not found" {
		return nil, err
	} else if err != nil && err.Error() == "record not found" {
		return nil, nil
	}
	return &user, nil
}

func (db *DB) UpdateUserById(u_id uuid.UUID, user *models.Users) error {
	err := db.db_conn.Model(&user).Where("id = ?", u_id).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) DeleteUserById(u_id uuid.UUID) error {
	err := db.db_conn.Where("id = ?", u_id).Update("is_deleted", true).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetAllUsers() ([]*models.Users, error) {
	var users []*models.Users
	err := db.db_conn.Where("is_deleted = ?", false).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (db *DB) GetUserByNumber(phone_number string) (*models.Users, error) {
	var user *models.Users
	err := db.db_conn.Where("phone_number = ?", phone_number).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DB) CheckUserByNumber(phone_number string) (bool, error) {
	var user models.Users
	err := db.db_conn.Where("phone_number = ? AND is_deleted = ?", phone_number, false).First(&user).Error
	if err != nil && err.Error() == "record not found" {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
