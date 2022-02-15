package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goginsoap/service"
	"goginsoap/soapHandler"
	"net/http"
)

type Response struct {
	Message string
	Error   string
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get productId
	r.GET("/api/product/:id/stock", func(c *gin.Context) {
		productId := c.Param("id")
		sede := c.Query("sede")
		sociedad := c.Query("sociedad")

		var request soapHandler.Request
		request.CodigoProducto = productId
		request.CodigoSede = sede
		request.CodigoSociedad = sociedad

		response, err := service.RetrieveStock(request)
		if err != nil {
			panic(err)
		}

		if err := json.NewEncoder(c.Writer).Encode(response); err != nil {
			panic(err)
		}
	})

	return r
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
