package handlers

import (
	"database/sql"
    "net/http"
	"strconv"
	
	"go-echo-vue/models"

    "github.com/labstack/echo"
)

type H map[string]interface{}

//GET Tasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error{
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// PutTasks endpoint
func PutTask(db *sql.DB) echo.HandlerFunc{
	return func(c echo.Context) error{
		// instanciar una nueva task
		var task models.Task
		c.Bind(&task)
		id, err := models.PutTask(db, task.Name)
		if err == nil{
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else{
			return err
		}
		
	}
}

// Delete task endpoint

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error{
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeleteTask(db, id)
		if err == nil{
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
		
	}
}