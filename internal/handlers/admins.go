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

type adminRequest struct {
	FullName string `json:"fullName" binding:"required,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func (h *Handlers) GetAllAdmins(c *gin.Context) {
	var admins []models.Admin
	if err := h.DB.Find(&admins).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"admins": admins,
	})
}

func (h *Handlers) CreateAdmin(c *gin.Context) {
	validatorError := make(map[string]string)
	var requestBody adminRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validatorError)
	}

	if err := h.DB.Where("email = ?", requestBody.Email).First(&models.Admin{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if _, exists := validatorError["email"]; !exists {
			validatorError["email"] = helpers.ValidationMessageForTag("unique", "")
		}
	}

	if len(validatorError) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validatorError)
		return
	}

	admin := models.Admin{
		FullName: requestBody.FullName,
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}
	if err := h.DB.Create(&admin).Error; err != nil {
		log.Println("Error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"admin": admin,
	})
}

func (h *Handlers) GetOneAdmin(c *gin.Context) {
	var admin models.Admin

	if err := h.DB.First(&admin, "id = ?", c.Param("id")).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Record not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"admin": admin,
	})
}

func (h *Handlers) UpdateAdmin(c *gin.Context) {
	var admin models.Admin
	if err := h.DB.First(&admin, "id = ?", c.Param("id")).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Record not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
		}
		return
	}

	validationErrors := make(map[string]string)
	var requestBody adminRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		helpers.FillValidationErrorTag(err, validationErrors)
	}

	if err := h.DB.Where("email = ?", requestBody.Email).
		Where("id != ?", admin.ID).
		First(&models.Admin{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if _, exists := validationErrors["email"]; !exists {
			validationErrors["email"] = helpers.ValidationMessageForTag("unique", "")
		}
	}

	if len(validationErrors) != 0 {
		c.JSON(http.StatusUnprocessableEntity, validationErrors)
		return
	}

	admin.FullName = requestBody.FullName
	admin.Email = requestBody.Email
	admin.Password = requestBody.Password

	if err := h.DB.Save(&admin).Error; err != nil {
		log.Println("cannot update admin:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"admin": admin,
	})
}

func (h *Handlers) DeleteAdmin(c *gin.Context) {
	var admin models.Admin

	if err := h.DB.First(&admin, "id = ?", c.Param("id")).Error; err != nil {
		log.Println("Error:", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Record not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
		}
		return
	}

	if err := h.DB.Delete(&admin).Error; err != nil {
		log.Println("Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Admin deleted successfully",
	})
}
