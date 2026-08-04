package travel

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luizfelipeeoliveiraac/travel-wallter-back-end/internal/auth"
)

type TravelHandler struct {
	repo *TravelRepository
}

func NewTravelHandler(repo *TravelRepository) *TravelHandler {
	return &TravelHandler{repo: repo}
}

func (h *TravelHandler) CreateTravel(c *gin.Context) {
	userID := getUserID(c)
	var req auth.CreateTravelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	travel, err := h.repo.CreateTravel(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, travel)
}

func (h *TravelHandler) GetTravel(c *gin.Context) {
	userID := getUserID(c)
	travelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid travel id"})
		return
	}

	travel, err := h.repo.GetTravelByID(travelID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "travel not found"})
		return
	}

	c.JSON(http.StatusOK, travel)
}

func (h *TravelHandler) ListTravels(c *gin.Context) {
	userID := getUserID(c)

	travels, err := h.repo.ListTravels(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if travels == nil {
		travels = []auth.Travel{}
	}

	c.JSON(http.StatusOK, travels)
}

func (h *TravelHandler) UpdateTravel(c *gin.Context) {
	userID := getUserID(c)
	travelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid travel id"})
		return
	}

	var req auth.CreateTravelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	travel, err := h.repo.UpdateTravel(travelID, userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, travel)
}

func (h *TravelHandler) DeleteTravel(c *gin.Context) {
	userID := getUserID(c)
	travelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid travel id"})
		return
	}

	if err := h.repo.DeleteTravel(travelID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "travel deleted"})
}

func getUserID(c *gin.Context) int {
	return auth.GetUserID(c)
}
