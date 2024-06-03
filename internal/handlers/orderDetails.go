package handlers

import (
	"apiAcademy/internal/database/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func (h *Handlers) GetAllOrderDetails(c *gin.Context) {
	var orderDetails []models.OrderDetail

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offset := (page - 1) * pageSize

	if err := h.DB.Limit(pageSize).Offset(offset).Find(&orderDetails).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orderDetails": orderDetails,
	})
}

func (h *Handlers) CreateOrderDetail(c *gin.Context) {
	var orderDetail models.OrderDetail

	if err := c.ShouldBindJSON(&orderDetail); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if err := h.DB.Create(&orderDetail).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, orderDetail)
}

func (h *Handlers) GetOneOrderDetail(c *gin.Context) {
	id := c.Param("id")
	var orderDetail models.OrderDetail

	if err := h.DB.First(&orderDetail, id).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "OrderDetail not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	c.JSON(http.StatusOK, orderDetail)
}

func (h *Handlers) UpdateOrderDetail(c *gin.Context) {
	id := c.Param("id")
	var orderDetail models.OrderDetail

	if err := h.DB.First(&orderDetail, id).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "OrderDetail not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	if err := c.ShouldBindJSON(&orderDetail); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if err := h.DB.Save(&orderDetail).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, orderDetail)
}

func (h *Handlers) DeleteOrderDetail(c *gin.Context) {
	id := c.Param("id")
	var orderDetail models.OrderDetail

	if err := h.DB.First(&orderDetail, id).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "OrderDetail not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	if err := h.DB.Delete(&orderDetail).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OrderDetail deleted successfully",
	})
}
