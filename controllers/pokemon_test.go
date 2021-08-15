package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPokemons(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", "/api/v1/pokemons", nil)
	if err != nil {
		panic(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200, "response code should equal 200")
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api/v1/")
	{
		grp1.GET("pokemons/", GetPokemons)
	}
	return r
}
