package controllers

import (
	"net/http"
	"sd_concepts/url_shortener/internal/services"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	svc *services.Service
}

func NewController(svc *services.Service) *Controller {
	return &Controller{
		svc: svc,
	}
}

func (ctl *Controller) GetHandler(c *gin.Context) {
	hash := c.Param("hash")

	longURL, err := ctl.svc.GetUrlService(c.Request.Context(), hash)

	if err != nil {
		code := http.StatusInternalServerError
		if err.Error() == "URL not found" {
			code = http.StatusNotFound
		}

		c.JSON(code, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"long_url": longURL,
	})
}

func (ctl *Controller) PostHandler(c *gin.Context) {
	var req struct {
		URL string `json:"url" binding:"required,url"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	hash, err := ctl.svc.CreateURLService(c.Request.Context(), req.URL)

	if err != nil {
		code := http.StatusInternalServerError
		if err.Error() == "already exists" {
			code = http.StatusConflict
		}

		c.JSON(code, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"short_url": hash,
	})
}

func (ctl *Controller) RedirectHandler(c *gin.Context) {
	hash := c.Param("hash")

	longURL, err := ctl.svc.GetUrlService(c.Request.Context(), hash)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "URL not found",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, longURL)
}
