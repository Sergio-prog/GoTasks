package api

import (
	"GoTasks/ToFile"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskDataFormat struct {
	Text string `json:"text"`
}

var file ToFile.File = ToFile.File{Path: "tasks.json"}

func getTasks(c *gin.Context) {
	data, err := file.GetData()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":     false,
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"result": string(data),
		"error":  nil,
	})
}

func postTasks(c *gin.Context) {
	var newTask TaskDataFormat

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":     false,
			"result": nil,
			"error":  err,
		})
		return
	}

	jsonData, err := json.Marshal(newTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":     false,
			"result": nil,
			"error":  err,
		})
		return
	}

	if err := file.AddData(string(jsonData)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":     false,
			"result": nil,
			"error":  err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ok":     true,
		"result": string(jsonData),
		"error":  nil,
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/tasks", getTasks)
	r.POST("/tasks", postTasks)
	return r
}

func Run() {
	r := setupRouter()
	r.Run(":8080")
}
