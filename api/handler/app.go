package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yeluoshaqi/config/api/config"
)

func CreateApp(c *gin.Context) {
	var req = struct {
		AppName     string `json:"appName"        binding:"required"`
		Description string `json:"description"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	app, err := config.CreateApp(req.AppName, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

func QueryApp(c *gin.Context) {
	var req = struct {
		AppName string `form:"appName"        binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	app, err := config.QueryApp(req.AppName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

func DeleteApp(c *gin.Context) {
	var req = struct {
		AppName string `form:"appName"        binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	err := config.DeleteApp(req.AppName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func ListApps(c *gin.Context) {
	apps, err := config.ListApps()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	if apps == nil {
		c.JSON(http.StatusOK, []interface{}{})
	} else {
		c.JSON(http.StatusOK, apps)
	}
}
