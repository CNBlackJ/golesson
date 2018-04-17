package hello

import (
	"time"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type News struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	Author          string `json:"author"`
	CreateAt        time.Time `json:"createAt"`
	UpdateAt        time.Time `json:"updateAt"`
}

var now = time.Now()
var news = [...] News {
	{
		Id: "0",
		Title: "Go lang study plan",
		Content: "Good Good Study, Day Day Up!!",
		Author: "Vinli Cheung",
		CreateAt: now,
		UpdateAt: now,
	},
	{
		Id: "1",
		Title: "Gin framwork",
		Content: "It is very easy to learn!",
		Author: "Vinli Cheung",
		CreateAt: now,
		UpdateAt: now,
	},
}

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if ( err != nil) {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"data": news[id],
		"code": 200,
		"serverTime": now,
	})
}

// func List(c *gin.Context) {
// 	limit := c.DefaultQuery("limit", 10)
// 	skip := strconv.Atoi(c.Query("skip"))
// 	c.JSON(200, gin.H{
// 		"data": news,
// 		"code": 200,
// 		"serverTime": now,
// 	})
// }

func Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "I am Post",
	})
}

func Put(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "I am Put",
	})
}

func Destroy(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "I am Delete",
	})
}