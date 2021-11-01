package controller

import (
	"net/http"

	"github.com/Chanon1359/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /ambulances
func CreateAmbulance(c *gin.Context) {
	var ambulance entity.Ambulance
	if err := c.ShouldBindJSON(&ambulance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&ambulance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulance})
}

// GET /ambulance/:id
func GetAmbulance(c *gin.Context) {
	var ambulance entity.Ambulance
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM ambulances WHERE id = ?", id).Find(&ambulance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ambulance})
}

// GET /ambulances
func ListAmbulances(c *gin.Context) {
	var ambulances []entity.Ambulance
	if err := entity.DB().Raw("SELECT * FROM ambulances").Find(&ambulances).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ambulances})
}

// DELETE /ambulances/:id
func DeleteAmbulance(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM ambulances WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulance not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /ambulances
func UpdateAmbulance(c *gin.Context) {
	var ambulance entity.Ambulance
	if err := c.ShouldBindJSON(&ambulance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", ambulance.ID).First(&ambulance); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulance not found"})
		return
	}

	if err := entity.DB().Save(&ambulance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ambulance})
}