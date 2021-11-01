package controller

import (
	"net/http"

	"github.com/Topzzson/SA-AM/entity"
	"github.com/gin-gonic/gin"
)

// POST /notifys
func CreateNotify(c *gin.Context) {
	var notify entity.Notify
	if err := c.ShouldBindJSON(&notify); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&notify).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": notify})
}

// GET /notify/:id
func GetNotify(c *gin.Context) {
	var notify entity.Notify

	id := c.Param("id")
	if err := entity.DB().Preload("Officer").Raw("SELECT * FROM notifies WHERE id = ?", id).Find(&notify).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notify})
}

// GET /notifys
func ListNotifys(c *gin.Context) {
	var notifys []entity.Notify
	if err := entity.DB().Preload("Officer").Raw("SELECT * FROM notifies").Find(&notifys).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notifys})
}

// DELETE /notifys/:id
func DeleteNotify(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM notifies WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "notify not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /notifys
func UpdateNotify(c *gin.Context) {
	var notify entity.Notify
	if err := c.ShouldBindJSON(&notify); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", notify.ID).First(&notify); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "notifies not found"})
		return
	}

	if err := entity.DB().Save(&notify).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notify})
}
