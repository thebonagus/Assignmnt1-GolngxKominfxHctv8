package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"strconv"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateSocmed(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Socialmedia := models.Socialmedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Socialmedia)
	} else {
		c.ShouldBind(&Socialmedia)
	}

	Socialmedia.UserID = userID

	if Socialmedia.UserID == userID {
		err := db.Debug().Create(&Socialmedia).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               Socialmedia.ID,
		"name":             Socialmedia.Name,
		"social_media_url": Socialmedia.SocialmediaURL,
		"user_id":          Socialmedia.UserID,
		"created_at":       Socialmedia.CreatedAt,
	})
}

func GetSocmed(c *gin.Context) {
	db := database.GetDB()
	Socialmedia := []models.Socialmedia{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	userID := uint(userData["id"].(float64))

	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Email", "Username")
	}).Where("user_id = ?", userID).Find(&Socialmedia).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Socialmedia)
}

func UpdateSocmed(c *gin.Context) {
	db := database.GetDB()
	Socialmedia := models.Socialmedia{}
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	socialmediaID, err := strconv.Atoi(c.Param("socialmediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "failed to convert",
		})
	}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socialmedia)
	} else {
		c.ShouldBind(&Socialmedia)
	}

	Socialmedia.UserID = userID
	Socialmedia.ID = uint(socialmediaID)

	err = db.Debug().Where("id=?", socialmediaID).Updates(models.Socialmedia{Name: Socialmedia.Name, SocialmediaURL: Socialmedia.SocialmediaURL}).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               Socialmedia.ID,
		"name":             Socialmedia.Name,
		"social_media_url": Socialmedia.SocialmediaURL,
		"user_id":          Socialmedia.UserID,
		"updated_at":       Socialmedia.UpdatedAt,
	})
}

func DeleteSocmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Socialmedia := models.Socialmedia{}

	socialmediaID, err := strconv.Atoi(c.Param("socialmediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "failed to convert",
		})
	}

	userID := uint(userData["id"].(float64))
	Socialmedia.UserID = userID
	Socialmedia.ID = uint(socialmediaID)

	err = db.Debug().Where("id = ?", socialmediaID).Delete(&Socialmedia).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
