package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/voje/gintrace"
)

func handleFuncPutHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
		"param":   42,
	})
}

func main() {
	fmt.Println("Set env TRACE_REST_BODY to get detailed output and slower performance.")

	log.SetLevel(log.TraceLevel)

	// We can also set this via ENV variables
	// gin.SetMode(gin.DebugMode) // default
	gin.SetMode(gin.ReleaseMode)

	r := gin.New() // Without middleware attached
	// r := gin.Default() // With some basic logging attached

	// Register middleware (this is where we modify our code).
	if log.GetLevel() == log.TraceLevel {
		r.Use(gintrace.BodyLogger)
	}

	// Test with
	// $ curl -X PUT -d arg=val -d arg2=val2 localhost:8080/hello
	r.PUT("/hello", handleFuncPutHello)

	r.Run("localhost:8080")
}
