package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pliniogsnascimento/helm-values-api/pkg/utils"
)

func main() {
	port := "7070"
	r := gin.Default()

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

	r.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
