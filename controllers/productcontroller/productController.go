package productcontrollers

import (
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mfhmiii/go-rest-api-1/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	//membuat slices
	var products []models.Product

	//mengambil semua data product di db
	models.DB.Find(&products)

	//kirim ke json
	c.JSON(http.StatusOK, gin.H{"product": products})
	//gin.H = maps{}
}

func Show(c *gin.Context) {
	//membuat slices
	var products models.Product

	//ngambil id di parameter routes /:id
	id := c.Param("id")

	//check if error with different type
	if err := models.DB.First(&products, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	//show the data
	c.JSON(http.StatusOK, gin.H{"product": products})
}

func Create(c *gin.Context) {
	var products models.Product

	//check if error
	if err := c.ShouldBindJSON(&products); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//create the product data in Database
	models.DB.Create(&products)

	//show the created Data
	c.JSON(http.StatusOK, gin.H{"product": products})
}

func Update(c *gin.Context) {
	var products models.Product
	// ngambil parameter di routes /:id
	id := c.Param("id")

	//check if error
	if err := c.ShouldBindJSON(&products); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//check if error, ngambil data dengan spesifik id + update data
	if models.DB.Model(&products).Where("id = ?", id).Updates(&products).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot Update Product"})
		return
	}

	//show the message
	c.JSON(http.StatusOK, gin.H{"message": "Data Updated Successfully"})
}

func Delete(c *gin.Context) {
    var input struct {
        ID json.Number `json:"id"`
    }

    // Check if there's an error in binding JSON data
    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    // Convert the ID to integer
    id, err := input.ID.Int64()
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
        return
    }

    // Delete the data
    if rowsAffected := models.DB.Delete(&models.Product{}, id).RowsAffected; rowsAffected == 0 {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot Delete Product"})
        return
    }

    // Show the success message
    c.JSON(http.StatusOK, gin.H{"message": "Data Deleted Successfully"})
}

