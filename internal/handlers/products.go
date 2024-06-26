package handlers

import (
	"apiAcademy/internal/database/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type productResponse struct {
	Name              string `json:"name" binding:"required" fake:"{name}"`
	Description       string `json:"description" binding:"required" fake:"{sentence}"`
	Price             uint   `json:"price" binding:"required" fake:"{number:1,100}"`
	QuantityAvailable int    `json:"quantity_available" binding:"required" fake:"{number:1,100}"`
	CategoryID        uint   `json:"category_id" binding:"required" fake:"{number: 1,10}"`
}

func (h *Handlers) GetAllProducts(c *gin.Context) {
	products := []models.Product{}
	if err := h.DB.Find(&products).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (h *Handlers) CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if err := h.DB.Create(&product).Error; err != nil {
		log.Println("Error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handlers) GetOneProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := h.DB.Preload("Reviews").First(&product, id).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Product not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	c.JSON(http.StatusOK, product)

}

func (h *Handlers) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := h.DB.First(&product, id).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Product not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if err := h.DB.Save(&product).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, product)

}

func (h *Handlers) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := h.DB.First(&product, id).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Product not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	if err := h.DB.Delete(&product).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
