package api

import (
	"GoTasks/ToFile"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func getTask(c *gin.Context) {
	intId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":     false,
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	data, err := file.GetData()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":     false,
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	var decodedData []map[string]string

	err = file.SafeJsonUnmarshal(data, &decodedData)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":     false,
			"result": nil,
			"error":  err.Error(),
		})
		return
	} else if intId >= len(decodedData) {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":     false,
			"result": nil,
			"error":  fmt.Sprintf("ID out of range of tasks. Len of task is %v.", len(decodedData)),
		})
		return
	}

	// result, err := json.Marshal(result_id)
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"ok":     false,
	// 		"result": nil,
	// 		"error":  err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"result": decodedData[intId],
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

	if newTask.Text == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":     false,
			"result": nil,
			"error":  "Invalid task format.",
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

func deleteTask(c *gin.Context) {
	fmt.Println("test")
	fmt.Println(c.Param("id"))
	intId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":     false,
			"result": nil,
			"error":  err.Error(),
		})
		return
	}

	if err := file.DeleteData(intId); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":     false,
			"result": nil,
			"error":  err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ok":     true,
		"result": true,
		"error":  nil,
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/tasks", getTasks)
	r.GET("/tasks/:id", getTask)
	r.POST("/tasks", postTasks)
	r.DELETE("/tasks/:id", deleteTask)
	return r
}

func Run() {
	r := setupRouter()
	r.Run(":8080")
}
