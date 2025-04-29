package service

import (
	"fmt"
	"housing_panda/models"
	"housing_panda/utils"

	"github.com/google/uuid"
)

func (s *Service) InsertListing(req *models.ListingRequest) error {

	err := utils.ValidateRequest(req)
	if err != nil {
		return fmt.Errorf("error validating request: %w", err)
	}

	listings := utils.RequestToListing(req)
	listings.Id = uuid.New()

	exists, err := s.db.CheckUserByNumber(req.PhoneNumber)
	if err != nil {
		return fmt.Errorf("error checking user by number: %w", err)
	}

	if exists {
		user, err := s.db.GetUserByNumber(req.PhoneNumber)
		if err != nil {
			return fmt.Errorf("error getting user by number: %w", err)
		}
		listings.UserId = user.Id
	} else {
		user := &models.Users{
			Id:          uuid.New(),
			PhoneNumber: req.PhoneNumber,
			Name:        req.Name,
		}
		err = s.db.InsertUser(user)
		if err != nil {
			return fmt.Errorf("error inserting user: %w", err)
		}
		listings.UserId = user.Id
	}

	err = s.db.InsertListing(listings)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetListings() ([]*models.ListingsResponse, error) {
	listings, err := s.db.GetAllListings()
	if err != nil {
		return nil, fmt.Errorf("error getting listings: %w", err)
	}

	var resListings []*models.ListingsResponse
	for _, listing := range listings {
		user, err := s.db.GetUserById(listing.UserId)
		if err != nil {
			return nil, fmt.Errorf("error getting user by id: %w", err)
		}
		resListing := utils.ListingToResponse(listing, user)
		resListings = append(resListings, resListing)
	}

	return resListings, nil
}

func (s *Service) DeleteListing(id string) error {
	listing_id, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("error parsing listing id: %w", err)
	}
	err = s.db.DeleteListingById(listing_id)
	if err != nil {
		return fmt.Errorf("error deleting listing: %w", err)
	}
	return nil
}

func (s *Service) GetListingById(id string) (*models.ListingsResponse, error) {
	listingId, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("error parsing listing id: %w", err)
	}
	listing, err := s.db.GetListingById(listingId)
	if err != nil {
		return nil, fmt.Errorf("error getting listing by id: %w", err)
	}
	if listing == nil {
		return nil, nil
	}
	user, err := s.db.GetUserById(listing.UserId)
	if err != nil {
		return nil, fmt.Errorf("error getting user by id: %w", err)
	}
	resListing := utils.ListingToResponse(listing, user)
	return resListing, nil
}

func (s *Service) GetListingsByUserId(id string) ([]*models.ListingsResponse, error) {
	user_id, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("error parsing user id: %w", err)
	}
	
	listings, err := s.db.GetAllListingsByUserId(user_id)
	if err != nil {
		return nil, fmt.Errorf("error getting listings by user id: %w", err)
	}
	if listings == nil {
		return nil, nil
	}

	var resListings []*models.ListingsResponse
	for _, listing := range listings {
		user, err := s.db.GetUserById(listing.UserId)
		if err != nil {
			return nil, fmt.Errorf("error getting user by id: %w", err)
		}
		resListing := utils.ListingToResponse(listing, user)
		resListings = append(resListings, resListing)
	}

	return resListings, nil
}
