package main

import "fmt"

// import "io/ioutil"
import "net/http"
import "github.com/gin-gonic/gin"

type wifi struct {
	Id   int
	Name string
}

type E struct {
	id   int
	name string
}

type FOO struct {
	URL string `json:"url" binding:"required"`
}

func listWifi(c *gin.Context) {
	wifis := [...]wifi{
		{
			Id:   1,
			Name: "tenda_wifi",
		},
		{
			Id:   2,
			Name: "xiaomi_wifi",
		},
	}

	wifi := wifi{
		Id:   1,
		Name: "a",
	}
	c.JSON(200, gin.H{
		"message": wifi,
	})
}

type NAME struct {
	NAME string `json:"name" binding:"required"`
}

func createWifi(c *gin.Context) {
	var name NAME
	c.BindJSON(&name)
	fmt.Printf("NAME to store: %v\n", name)
	c.JSON(http.StatusOK, name)
}

func getWifi(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": id,
	})
}

func main() {
	router := gin.Default()

	router.GET("/wifis", listWifi)

	router.POST("/wifis", createWifi)

	router.GET("/wifis/:id", getWifi)

	// router.PUT("/wifis/:id", updateWifi)

	// router.DELETE("/wifis/:id", deleteWifi)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/foo", func(c *gin.Context) {
		var url FOO
		c.BindJSON(&url)
		fmt.Printf("URL to store: %v\n", url)
		c.JSON(http.StatusOK, url)
	})

	router.Run()
}
