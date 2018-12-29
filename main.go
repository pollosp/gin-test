package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bucketsjson struct {
	Type string   `json:"type"`
	List []string `json:"list"`
}

func testData() bucketsjson {
	jsonSrc := []byte(`{
		"type":"bucket",
		"list":[
			 "canguro",
			 "ardilla"
		]
	}`)

	var buckets bucketsjson
	json.Unmarshal(jsonSrc, &buckets)
	return buckets
}

//SetupRouter configures the main router
func SetupRouter() *gin.Engine {
	router := gin.Default()
	buckets := testData()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, buckets)
	})

	return router
}

func main() {
	router := SetupRouter()
	router.Run()
}
