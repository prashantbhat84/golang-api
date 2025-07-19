package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Task represents a task with its properties.
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}
type TaskNoID struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

// Mock data for tasks
var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func main() {
	fmt.Println("Golang API")
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
	router.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, tasks)
	})
	router.GET("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, task := range tasks {
			if task.ID == id {
				resp := TaskNoID{
					Title:       task.Title,
					Description: task.Description,
					DueDate:     task.DueDate,
					Status:      task.Status,
				}
				c.JSON(http.StatusOK, resp)
				return
			}
		}
		c.JSON(404, gin.H{"message": "Task not found"})
	})
	router.PUT("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedTask TaskNoID
		if err := c.ShouldBindJSON(&updatedTask); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		// Print the parsed request body object
		fmt.Printf("Parsed Request Body: %+v\n", updatedTask)

		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Title = updatedTask.Title
				tasks[i].Description = updatedTask.Description
				tasks[i].DueDate = updatedTask.DueDate
				tasks[i].Status = updatedTask.Status
				c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	})

	router.Run(":8080")

}
