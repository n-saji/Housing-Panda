package utils

import (
	"fmt"
	"housing_panda/models"
)

func ValidateRequest(req *models.ListingRequest) error {

	if req.PhoneNumber == "" {
		return fmt.Errorf("phone number is required")
	}
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	if req.Rent <= 0 {
		return fmt.Errorf("rent must be greater than zero")
	}
	if req.NumberOfBedrooms <= 0 {
		return fmt.Errorf("number of bedrooms must be greater than zero")
	}
	if req.NumberOfBathrooms <= 0 {
		return fmt.Errorf("number of bathrooms must be greater than zero")
	}
	if req.Address == "" {
		return fmt.Errorf("address is required")
	}
	if req.Title == "" {
		return fmt.Errorf("title is required")
	}

	return nil
}
