package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pliniogsnascimento/helm-values-api/pkg/utils"
)

func main() {
	port := "7070"
	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health"),
		gin.Recovery(),
	)

	r.GET("/charts/:name/values", func(ctx *gin.Context) {
		name := ctx.Param("name")
		status, out := utils.RunHelmGetValues(name)
		ctx.Data(status, "application/yaml", out)
	})

	r.GET("/charts/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		status, out := utils.RunHelmGetValues(name)
		ctx.Data(status, "application/yaml", out)
	})

	r.GET("/charts", func(ctx *gin.Context) {
		release := ctx.Query("namespace")
		status, out := utils.RunHelmGetReleases(release)
		ctx.Data(status, "application/yaml", out)
	})

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Healthy")
	})

	r.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
