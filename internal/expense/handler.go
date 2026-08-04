package expense

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luizfelipeeoliveiraac/travel-wallter-back-end/internal/auth"
)

type ExpenseHandler struct {
	repo *ExpenseRepository
}

func NewExpenseHandler(repo *ExpenseRepository) *ExpenseHandler {
	return &ExpenseHandler{repo: repo}
}

func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	userID := getUserID(c)
	travelID, err := strconv.Atoi(c.Param("travel_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid travel id"})
		return
	}

	var req auth.CreateGastoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gasto, err := h.repo.CreateExpense(travelID, userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gasto)
}

func (h *ExpenseHandler) GetExpense(c *gin.Context) {
	userID := getUserID(c)
	travelID, err := strconv.Atoi(c.Param("travel_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid travel id"})
		return
	}

	gastoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expense id"})
		return
	}

	gasto, err := h.repo.GetExpenseByID(gastoID, travelID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "expense not found"})
		return
	}

	c.JSON(http.StatusOK, gasto)
}

func (h *ExpenseHandler) ListExpenses(c *gin.Context) {
	userID := getUserID(c)
	travelID, err := strconv.Atoi(c.Param("travel_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid travel id"})
		return
	}

	gastos, err := h.repo.ListExpenses(travelID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if gastos == nil {
		gastos = []auth.Gasto{}
	}

	c.JSON(http.StatusOK, gastos)
}

func (h *ExpenseHandler) UpdateExpense(c *gin.Context) {
	userID := getUserID(c)
	travelID, err := strconv.Atoi(c.Param("travel_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid travel id"})
		return
	}

	gastoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expense id"})
		return
	}

	var req auth.CreateGastoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gasto, err := h.repo.UpdateExpense(gastoID, travelID, userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gasto)
}

func (h *ExpenseHandler) DeleteExpense(c *gin.Context) {
	userID := getUserID(c)
	travelID, err := strconv.Atoi(c.Param("travel_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid travel id"})
		return
	}

	gastoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expense id"})
		return
	}

	if err := h.repo.DeleteExpense(gastoID, travelID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "expense deleted"})
}

func (h *ExpenseHandler) GetTravelTotal(c *gin.Context) {
	userID := getUserID(c)
	travelID, err := strconv.Atoi(c.Param("travel_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid travel id"})
		return
	}

	total, err := h.repo.GetTravelTotal(travelID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": total})
}

func getUserID(c *gin.Context) int {
	return auth.GetUserID(c)
}
