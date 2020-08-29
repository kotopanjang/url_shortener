package controllers

import (
	"aqilliz_assesment/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// Get raw query. behind the ? sign
	host := c.Request.Host
	oriUrl := c.Request.URL.RawQuery
	short := models.ShortURL{}
	short.OriginalUrl = oriUrl

	//give the stored url
	existingData := short.CheckExistingData()
	if existingData != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":    host + "/retrieve?" + existingData.MaskedUrl,
			"error":   false,
			"message": nil,
		})
	} else {
		data, err := short.Insert()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"data":    nil,
				"error":   true,
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data":    host + "/retrieve?" + data.MaskedUrl,
				"error":   false,
				"message": nil,
			})
		}

	}
}

func Retrieve(c *gin.Context) {
	// Get raw query. behind the ? sign
	query := c.Request.URL.RawQuery
	short := models.ShortURL{}
	res, err := short.Retrieve(query)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":    "",
			"error":   true,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    res,
			"error":   false,
			"message": nil,
		})
	}
}

func Redirect(c *gin.Context) {
	// Get raw query. behind the ? sign
	query := c.Request.URL.RawQuery
	short := models.ShortURL{}
	res, err := short.Retrieve(query)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":    "",
			"error":   true,
			"message": err.Error(),
		})
	} else {
		c.Redirect(301, res)
	}
}
