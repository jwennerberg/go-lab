package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/prometheus/client_golang/prometheus/promhttp"
  "github.com/common-nighthawk/go-figure"
)

func prometheusHandler() gin.HandlerFunc {
  h := promhttp.Handler()

  return func(c *gin.Context) {
    h.ServeHTTP(c.Writer, c.Request)
  }
}

func setupRouter() *gin.Engine {
  r := gin.Default()

  r.GET("/healthz", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"health": "ok"})
  })

  r.GET("/hello", func(c *gin.Context) {
    myFigure := figure.NewFigure("Hi", "rectangles", true)
    figure.Write(c.Writer, myFigure)
  })

  r.GET("/hello/:name", func(c *gin.Context) {
    name := c.Param("name")
    myFigure := figure.NewFigure("Hello " + name, "rectangles", true)
    figure.Write(c.Writer, myFigure)
    //c.String(http.StatusOK, "Hello %s\n", name)
  })

  r.GET("/metrics", prometheusHandler())

  return r
}

func main() {
  r := setupRouter()
  r.Run(":8080")
}
