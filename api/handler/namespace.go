package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yeluoshaqi/configuration/api/config"
)

func CreateNamespace(c *gin.Context) {
	var req = struct {
		AppName       string `json:"appName"        binding:"required"`
		ClusterName   string `json:"clusterName"    binding:"required"`
		NamespaceName string `json:"namespaceName"  binding:"required"`
		Format        string `json:"format"         binding:"required"`
		Description   string `json:"description"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	namespaces, err := config.CreateNamespace(req.AppName, req.ClusterName, req.NamespaceName, req.Format, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, namespaces)
}

func QueryNamespace(c *gin.Context) {
	var req = struct {
		AppName       string `form:"appName"        binding:"required"`
		ClusterName   string `form:"clusterName"    binding:"required"`
		NamespaceName string `form:"namespaceName"  binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	namespaces, err := config.QueryNamespace(req.AppName, req.ClusterName, req.NamespaceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, namespaces)
}

func DeleteNamespace(c *gin.Context) {
	var req = struct {
		AppName       string `form:"appName"        binding:"required"`
		ClusterName   string `form:"clusterName"    binding:"required"`
		NamespaceName string `form:"namespaceName"  binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	err := config.DeleteNamespace(req.AppName, req.ClusterName, req.NamespaceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func ListNamespaces(c *gin.Context) {
	var req = struct {
		AppName     string `form:"appName"        binding:"required"`
		ClusterName string `form:"clusterName"    binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	namespaces, err := config.ListNamespaces(req.AppName, req.ClusterName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	if namespaces == nil {
		c.JSON(http.StatusOK, []interface{}{})
	} else {
		c.JSON(http.StatusOK, namespaces)
	}
}
