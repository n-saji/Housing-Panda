package models

import "github.com/google/uuid"

type ListingRequest struct {
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	Address           string  `json:"address"`
	Rent              float64 `json:"rent"`
	NumberOfBedrooms  float64 `json:"number_of_bedrooms"`
	NumberOfBathrooms float64 `json:"number_of_bathrooms"`
	Name              string  `json:"name"`
	PhoneNumber       string  `json:"phone_number"`
}

type ListingsResponse struct {
	ListingId         uuid.UUID `json:"listing_id"`
	UserId            uuid.UUID `json:"user_id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Address           string    `json:"address"`
	Rent              float64   `json:"rent"`
	NumberOfBedrooms  float64   `json:"number_of_bedrooms"`
	NumberOfBathrooms float64   `json:"number_of_bathrooms"`
	Name              string    `json:"name"`
	PhoneNumber       string    `json:"phone_number"`
	CreatedAt         int64     `json:"created_at"`
	UpdatedAt         int64     `json:"updated_at"`
	IsDeleted         bool      `json:"is_deleted"`
}
