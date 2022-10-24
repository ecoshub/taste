package main

import (
	"net/http"

	example "github.com/ecoshub/taste/example/server"
	"github.com/gin-gonic/gin"
)

func exampleGINServer() *gin.Engine {
	e := gin.New()
	apiV1 := e.Group("/api/v1")
	apiV1.GET("version", versionHandlerGIN)
	apiV1.GET("users", usersHandlerGIN)
	apiV1.GET("user", userHandlerGIN)
	apiV1.POST("user/new", newUserHandlerGIN)
	return e
}

func versionHandlerGIN(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/plain", []byte("v1.0.0"))
}

func usersHandlerGIN(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, example.Users)
}

func userHandlerGIN(ctx *gin.Context) {
	name := ctx.Query("name")
	user, exists := example.GetUser(name)
	if !exists {
		ctx.Data(http.StatusNotFound, "text/plain", []byte("404 page not found"))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func newUserHandlerGIN(ctx *gin.Context) {
	user := &example.User{}
	err := ctx.BindJSON(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	example.AddUser(user)
	ctx.JSON(http.StatusOK, example.Users)
}
