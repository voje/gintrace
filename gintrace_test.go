package gintrace

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	router := gin.New()
	router.GET("/get123", func(c *gin.Context) {
		c.String(http.StatusOK, "get response")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get123", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, "get response", w.Body.String())
}

func TestPost(t *testing.T) {
	requestBody := []byte("{'requestKey':'requestVal'}")
	responseBody := []byte("{'responseKey':'requestVal'}")

	router := gin.New()

	router.POST("/post123", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": string(responseBody),
		})
	})

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/post123", bytes.NewBuffer(requestBody))

	router.ServeHTTP(w, req)

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, string(responseBody), body["message"])
}

func TestLogging(t *testing.T) {
	requestBody := []byte("{'requestKey':'requestVal'}")
	responseBody := []byte("{'responseKey':'requestVal'}")

	logFile := bytes.NewBuffer(nil)
	log.SetOutput(logFile)

	router := gin.New()

	// Set to trace and use BodyLogger
	log.SetLevel(log.TraceLevel)
	if log.GetLevel() == log.TraceLevel {
		router.Use(BodyLogger)
	}

	router.POST("/post123", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": string(responseBody),
		})
	})

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/post123", bytes.NewBuffer(requestBody))

	router.ServeHTTP(w, req)

	var body map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, string(responseBody), body["message"])

	scanner := bufio.NewScanner(logFile)

	scanner.Scan()
	loggedRequest := scanner.Text()
	assert.Contains(t, loggedRequest, "RequestBody: {'requestKey':'requestVal'}")

	scanner.Scan()
	loggedResponse := scanner.Text()
	assert.Contains(t, loggedResponse, `ResponseBody: {\"code\":200,\"message\":\"{'responseKey':'requestVal'}\"}`)
}
