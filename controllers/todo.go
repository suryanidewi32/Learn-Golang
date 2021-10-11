package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	guuid "github.com/google/uuid"
)

type Todo struct {
	ID        string    `json:"id"`
	//Title     string    `json:"username"`
	//Body      string    `json:"password"`
	Name 	  string    `json:"name"`
	Jabatan   string    `json:"jabatan"`
	Username  string    `json:"username"`
	Completed string    `json:"privilege"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DATABASE INSTANCE
var collection *mongo.Collection

func TodoCollection(c *mongo.Database) {
	collection = c.Collection("user_auth")
}

func GetAllTodos(c *gin.Context) {
	todos := []Todo{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Iterate through the returned cursor.
    for cursor.Next(context.TODO()) {
				var todo Todo
        cursor.Decode(&todo)
        todos = append(todos, todo)
		}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Todos",
		"data":    todos,
	})
	return
}

func CreateTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	name := todo.Name
	jabatan := todo.Jabatan
	username := todo.Username
	completed := todo.Completed
	id := guuid.New().String()

	newTodo := Todo{
		ID: id,
		Name:     name,
		Jabatan:      jabatan,
		Username: 		username,
		Completed: completed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := collection.InsertOne(context.TODO(), newTodo)

	if err != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo created Successfully",
	})
	return
}

func GetSingleTodo(c *gin.Context) {
	todoId := c.Param("todoId")

	todo := Todo{}
	err := collection.FindOne(context.TODO(), bson.M{"id": todoId}).Decode(&todo)
	if err != nil {
			log.Printf("Error while getting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Todo",
		"data": todo,
	})
	return
}

func EditTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	var todo Todo
	c.BindJSON(&todo)
	completed := todo.Completed

	newData := bson.M{
            "$set": bson.M{
            "completed":       completed,
            "updated_at": time.Now(),
            },
        }

	_, err := collection.UpdateOne(context.TODO(), bson.M{"id": todoId}, newData)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message":  "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Todo Edited Successfully",
	})
	return
}

func DeleteTodo(c *gin.Context) {
todoId := c.Param("todoId")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": todoId})
	if err != nil {
		log.Printf("Error while deleting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deleted successfully",
	})
	return
}
