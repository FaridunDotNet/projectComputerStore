package handlers

import (
	"apiAcademy/internal/database/models"
	"apiAcademy/internal/helpers"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type categoryRequest struct {
	Name string `json:"name" binding:"required,max=64"`
}

func (h *Handlers) GetAllCategories(c *gin.Context) {
	var categories []models.Category
	if err := h.DB.Find(&categories).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func (h *Handlers) CreateCategory(c *gin.Context) {
	validatorError := make(map[string]string)
	var requestBody categoryRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {

		helpers.FillValidationErrorTag(err, validatorError)
	}

	if len(validatorError) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validatorError)
		return
	}

	category := models.Category{
		Name: requestBody.Name,
	}
	if err := h.DB.Create(&category).Error; err != nil {
		log.Println("Error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"category": category,
	})
}

func (h *Handlers) GetOneCategory(c *gin.Context) {
	//id := c.Param("id")
	var category models.Category

	if err := h.DB.First("id = ?", c.Param("id")).
		First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Record not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})

}

func (h *Handlers) UpdateCategory(c *gin.Context) {
	var category models.Category
	if err := h.DB.Where("id = ?", c.Param("id")).
		First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Record Not Found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	validationErrors := make(map[string]string)

	var requestBody categoryRequest
	// data validation
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	// Business validation
	// select count(*) from categories where name='name' and id != 'id' // if count > 0

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	// Update category fields with requestBody data
	category.Name = requestBody.Name

	if err := h.DB.Save(&category).Error; err != nil {
		log.Println("cannot update category:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

func (h *Handlers) DeleteCategory(c *gin.Context) {
	var category models.Category

	if err := h.DB.First("id = ?", c.Param("id")).
		First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Record not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	if err := h.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully",
	})
}
