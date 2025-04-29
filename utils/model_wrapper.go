package utils

import "housing_panda/models"

func RequestToListing(req *models.ListingRequest) *models.Listings {
	return &models.Listings{
		Title:             req.Title,
		Description:       req.Description,
		Rent:              req.Rent,
		Address:           req.Address,
		NumberOfBedrooms:  req.NumberOfBedrooms,
		NumberOfBathrooms: req.NumberOfBathrooms,
	}
}

func ListingToResponse(listing *models.Listings, user *models.Users) *models.ListingsResponse {
	return &models.ListingsResponse{
		ListingId:         listing.Id,
		UserId:            user.Id,
		Title:             listing.Title,
		Description:       listing.Description,
		Rent:              listing.Rent,
		Address:           listing.Address,
		NumberOfBedrooms:  listing.NumberOfBedrooms,
		NumberOfBathrooms: listing.NumberOfBathrooms,
		Name:              user.Name,
		PhoneNumber:       user.PhoneNumber,
		CreatedAt:         listing.CreatedAt,
		UpdatedAt:         listing.UpdatedAt,
		IsDeleted:         listing.IsDeleted,
	}
}
