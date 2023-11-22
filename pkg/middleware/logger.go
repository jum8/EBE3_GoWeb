package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verb := ctx.Request.Method
		time := time.Now()
		url := ctx.Request.URL
		var size int
		
		ctx.Next()

		if ctx.Writer != nil {
			size = ctx.Writer.Size()
		}

		fmt.Printf("Verbo:%s\tPath:%s\tFecha y hora:%s\tTamaño consulta:%d\n", verb, url, time.Format("02-01-2006 15:04:05"), size)
	}
}