package controllers

import (
	"net/http"
	"strings"

	"github.com/kotopanjang/url_shortener/helper"
	"github.com/kotopanjang/url_shortener/models"

	"github.com/gin-gonic/gin"
)

// This function is to handle register handler
func Register(c *gin.Context) {
	// Get raw query. behind the ? sign
	host := c.Request.Host
	oriUrl := c.Request.URL.RawQuery
	// if strings.

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

// This function is to handle Retrieve handler
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

// This function is to handle Redirect handler
func Redirect(c *gin.Context) {
	// Get raw query. behind the ? sign
	query := c.Request.URL.RawQuery
	short := models.ShortURL{}
	res, err := short.Retrieve(query)
	helper.Println("NGAGNNUUUUUUU")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":    "",
			"error":   true,
			"message": err.Error(),
		})
	} else {
		helper.Println("Prefix >> ", !strings.HasPrefix(res, "http"))
		if !strings.HasPrefix(res, "http") {
			c.Redirect(301, res)
		} else {
			c.Redirect(301, "http://www."+res)
		}
	}
}
