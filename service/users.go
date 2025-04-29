package service

import (
	"fmt"
	"housing_panda/models"

	"github.com/google/uuid"
)

func (s *Service) InsertUsers(user *models.Users) error {

	// check if phone number is already registered
	exists, err := s.db.CheckUserByNumber(user.PhoneNumber)
	if err != nil {
		return fmt.Errorf("error checking user by number: %w", err)
	}

	if exists {
		return fmt.Errorf("user with phone number %s already exists", user.PhoneNumber)
	}

	if err := s.db.InsertUser(user); err != nil {
		return err
	}
	return nil

}

func (s *Service) GetUserById(u_id string) (*models.Users, error) {
	u_id_parsed, err := uuid.Parse(u_id)
	if err != nil {
		return nil, fmt.Errorf("error parsing user id: %w", err)
	}
	user, err := s.db.GetUserById(u_id_parsed)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) UpdateUserById(u_id string, user *models.Users) error {
	u_id_parsed, err := uuid.Parse(u_id)
	if err != nil {
		return fmt.Errorf("error parsing user id: %w", err)
	}
	user.Id = u_id_parsed

	err = s.db.UpdateUserById(u_id_parsed, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUserById(u_id string) error {
	u_id_parsed, err := uuid.Parse(u_id)
	if err != nil {
		return fmt.Errorf("error parsing user id: %w", err)
	}
	listings, err := s.db.GetAllListingsByUserId(u_id_parsed)
	if err != nil {
		return fmt.Errorf("error getting all listings: %w", err)
	}
	for _, listing := range listings {
		err = s.db.DeleteListingById(listing.Id)
		if err != nil {
			return fmt.Errorf("error deleting listing by id: %w", err)
		}
	}

	err = s.db.DeleteUserById(u_id_parsed)
	if err != nil {
		return fmt.Errorf("error deleting user by id: %w", err)
	}

	return nil
}

func (s *Service) GetAllUsers() ([]*models.Users, error) {
	users, err := s.db.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("error getting all users: %w", err)
	}
	return users, nil
}

func (s *Service) GetUserByNumber(phone_number string) (*models.Users, error) {
	user, err := s.db.GetUserByNumber(phone_number)
	if err != nil {
		return nil, fmt.Errorf("error getting user by number: %w", err)
	}
	return user, nil
}
