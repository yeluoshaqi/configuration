package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/yeluoshaqi/configuration/api/handler"
	"github.com/yeluoshaqi/configuration/config-srv/conf"
	"github.com/yeluoshaqi/configuration/config-srv/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	confi := &conf.Config{}
	confi.DB.DriverName = "mysql"
	confi.DB.URL = "root@(127.0.0.1:3306)/config?charset=utf8&parseTime=true&loc=Local"

	var err error
	if err = dao.Init(confi); err != nil {
		panic(err)
	}
	if err = dao.GetDao().Ping(); err != nil {
		panic(err)
	}
	defer  dao.GetDao().Disconnect()

	router := gin.Default()
	router.Static("/admin/ui", "./dist")
	r := router.Group("/admin/api/v1")
	r.GET("/apps", handler.ListApps)
	r.GET("/app", handler.QueryApp)
	r.POST("/app", handler.CreateApp)
	r.DELETE("/app", handler.DeleteApp)

	r.GET("/clusters", handler.ListClusters)
	r.GET("/cluster", handler.QueryCluster)
	r.POST("/cluster", handler.CreateCluster)
	r.DELETE("/cluster", handler.DeleteCluster)

	r.GET("/namespaces", handler.ListNamespaces)
	r.GET("/namespace", handler.QueryNamespace)
	r.POST("/namespace", handler.CreateNamespace)
	r.DELETE("/namespace", handler.DeleteNamespace)

	r.POST("/config", handler.UpdateConfig)
	r.POST("/release", handler.Release)
	r.POST("/rollback", handler.Rollback)
	r.GET("/release/history", handler.ListReleaseHistory)

	r.GET("/format", handler.ListSupportedFormat)

	router.Run(":8808")
}
