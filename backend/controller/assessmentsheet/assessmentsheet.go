package controller

import (
	"net/http"

	"github.com/Chanon1359/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /assessmentsheets
func CreateAssessmentSheet(c *gin.Context) {
	var assessmentsheet entity.AssessmentSheet
	if err := c.ShouldBindJSON(&assessmentsheet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&assessmentsheet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": assessmentsheet})
}

func GetAssessmentSheet(c *gin.Context) {
	var assessmentsheet entity.AssessmentSheet
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM assessment_sheets WHERE id = ?", id).Find(&assessmentsheet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": assessmentsheet})
}

// GET /assessmentsheets
func ListAssessmentSheets(c *gin.Context) {
	var assessmentsheets []entity.AssessmentSheet
	if err := entity.DB().Raw("SELECT * FROM assessment_sheets").Find(&assessmentsheets).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": assessmentsheets})
}



// DELETE /assessmentsheets/:id
func DeleteAssessmentSheet(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM assessment_sheets WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assessmentsheets not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /assessmentsheets
func UpdateAssessmentSheet(c *gin.Context) {
	var assessmentsheet entity.AssessmentSheet
	if err := c.ShouldBindJSON(&assessmentsheet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", assessmentsheet.ID).First(&assessmentsheet); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assessment_sheets not found"})
		return
	}

	if err := entity.DB().Save(&assessmentsheet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": assessmentsheet})
}