package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	statusOK         = 200
	statusBadRequest = 400
	statusNotFound   = 404
)

// Payload query structure
type Payload struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// GQLHandler used by grahpql schemas
// Accepts data from Form Data and Request Payload
func GQLHandler(c *gin.Context) {
	payload := Payload{
		Query:     c.PostForm("query"),
		Variables: make(map[string]interface{}),
	}
	err := getPayload(c, &payload)
	if err != nil {
		c.JSON(statusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(statusOK, GraphQL(&payload))
}

func getPayload(c *gin.Context, payload *Payload) error {
	// check for existence of data from Form Data
	if payload.Query == "" {
		// if not, then we take from Request Payload
		if c.Request.Body == nil {
			return errors.New("Please send a Payload or Form Data")
		}
		err := json.NewDecoder(c.Request.Body).Decode(&payload)
		if err != nil {
			return err
		}
	}
	return nil
}

// OptionsHandler is needed for the front, since the method is first sent with OPTIONS
func OptionsHandler(c *gin.Context) {
	c.JSON(statusOK, "")
}

// SetHeaders sets headers
func SetHeaders(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept")
	c.Header("Access-Control-Max-Age", "600")
	c.Header("Connection", "keep-alive")
}
