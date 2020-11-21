package gintrace

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	router := gin.New()
	router.GET("/test1", func(c *gin.Context) {
		c.String(200, "test1r")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, "test1r", w.Body.String())
}
func TestPost(t *testing.T) {
	router := gin.New()
	router.POST("/test1", func(c *gin.Context) {
		retJSON := struct {
			a int
			b string
		}{
			42,
			"test",
		}
		c.JSON(200, retJSON)
	})

	w := httptest.NewRecorder()
	reqJSON := struct {
		a int
	}{
		1,
	}
	jsonBytes, _ := json.Marshal(reqJSON)
	br := bytes.NewBuffer(jsonBytes)
	req, _ := http.NewRequest("POST", "/test1", br)

	router.ServeHTTP(w, req)

	assert.Equal(t, "test1rrr", w.Body.String())
}
