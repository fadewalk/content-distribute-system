package api

import (
	"ContentSystem/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	rootPath   = "/api/"
	noAuthPath = "/out/api/"
)

func CmsRouters(r *gin.Engine) {
	cmsApp := services.NewCmsApp()
	session := NewSessionAuth()
	// 开启上报
	r.Use(prometheusMiddleware())
	//r.Use(traceMiddleware(r))
	r.Use(opentracingMiddleware())
	// 逻辑路由
	root := r.Group(rootPath).Use(session.Auth)
	{
		// /api/cms/hello
		root.GET("/cms/hello", cmsApp.Hello)
		// /api/cms/create
		root.POST("/cms/content/create", cmsApp.ContentCreate)
		// /api/cms/update
		root.POST("/cms/content/update", cmsApp.ContentUpdate)
		// /api/cms/delete
		root.POST("/cms/content/delete", cmsApp.ContentDelete)
		// /api/cms/find
		root.POST("/cms/content/find", cmsApp.ContentFind)
	}
	noAuth := r.Group(noAuthPath)
	{
		// /out/api/cms/register
		noAuth.POST("/cms/register", cmsApp.Register)
		// /out/api/cms/login
		noAuth.POST("/cms/login", cmsApp.Login)
	}
	// http://localhost:8080/metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
