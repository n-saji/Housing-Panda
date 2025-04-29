package handlers

import (
	"housing_panda/models"
	"housing_panda/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handlers struct {
	service *service.Service
}

func NewHandlers(db *gorm.DB) *Handlers {
	return &Handlers{
		service: service.NewService(db),
	}
}

func (h *Handlers) GetRouter() *gin.Engine {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5050"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin", "content-type", "Set-Cookie"},
		ExposeHeaders:    []string{"Content-Length", "Set-Cookie"},
		AllowCredentials: true,
	}))

	h.RoutingChannel(&router.RouterGroup)

	return router
}

func (h *Handlers) RoutingChannel(rc *gin.RouterGroup) {
	rc.POST("/listings", h.InsertListing)
	rc.GET("/listings", h.GetListings)
	rc.GET("/listings/:listing_id", h.GetListingById)
	rc.GET("/listings/user/:user_id", h.GetListingsByUserId)
	rc.DELETE("/listings/:listing_id", h.DeleteListing)
	rc.GET("/users/:user_id", h.GetUserById)
	rc.GET("/users", h.GetUsers)

	// rc.POST("/users", h.InsertUser)
	// rc.PUT("/listings/:listing_id", h.UpdateListing)
}

func (h *Handlers) InsertListing(c *gin.Context) {
	req := &models.ListingRequest{}
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	err = h.service.InsertListing(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "listing inserted successfully"})
}

func (h *Handlers) GetListings(c *gin.Context) {
	listings, err := h.service.GetListings()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, listings)
}

func (h *Handlers) GetListingById(c *gin.Context) {
	listing_id := c.Param("listing_id")
	listing, err := h.service.GetListingById(listing_id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, listing)
}

func (h *Handlers) DeleteListing(c *gin.Context) {
	listing_id := c.Param("listing_id")
	err := h.service.DeleteListing(listing_id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "deleted successfully"})
}

func (h *Handlers) GetUserById(c *gin.Context) {
	user_id := c.Param("user_id")
	user, err := h.service.GetUserById(user_id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (h *Handlers) GetUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

func (h *Handlers) GetListingsByUserId(c *gin.Context) {
	user_id := c.Param("user_id")
	listings, err := h.service.GetListingsByUserId(user_id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, listings)
}
