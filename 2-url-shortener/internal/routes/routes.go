package routes

import (
	"sd_concepts/url_shortener/internal/controllers"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, ctl *controllers.Controller) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/urls/:hash", ctl.GetHandler)
		v1.POST("/urls", ctl.PostHandler)
	}

	r.GET("/:hash", ctl.RedirectHandler)
}
