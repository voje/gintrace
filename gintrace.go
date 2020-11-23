package gintrace

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Implement our own response writer, so we can catch response body
type myResponseWriter struct {
	gin.ResponseWriter               // Using original ResponseWriter
	body               *bytes.Buffer // Adding save response buffer
}

func (mr *myResponseWriter) Write(respBody []byte) (int, error) {
	// Save bytes for later
	_, err := mr.body.Write(respBody)
	if err != nil {
		return 0, err
	}
	// Proceed with default behavior
	return mr.ResponseWriter.Write(respBody)
}

// BodyLogger logs both request and response body
func BodyLogger(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		log.Error(err)
	}
	fmt.Printf("RequestBody: %s\n", string(body))

	// We also want to catch response (it gets sent off during the next handler)
	// We can implement our own c.Writer that saves the response before sending it off
	w := &myResponseWriter{
		body:           &bytes.Buffer{},
		ResponseWriter: c.Writer,
	}
	c.Writer = w

	// Next handleFunc in chain
	c.Next()

	// Log the stored response
	fmt.Printf("ResponseBody: %s\n", w.body.String())
}
