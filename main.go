package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//go get -u github.com/labstack/echo/v4

type todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []todo{
	{ID: 1, Title: "Buy groceries", Done: false},
	{ID: 2, Title: "Do laundry", Done: true},
	{ID: 3, Title: "Clean room", Done: false},
}

func main() {
	e := echo.New()

	// GET /todos
	e.GET("/todos", func(c echo.Context) error {
		return c.JSON(http.StatusOK, todos)
	})

	// POST /todos
	e.POST("/todos", func(c echo.Context) error {
		var t todo
		if err := c.Bind(&t); err != nil {
			return err
		}
		t.ID = len(todos) + 1
		todos = append(todos, t)
		return c.JSON(http.StatusCreated, t)
	})

	// GET /todos/:id
	e.GET("/todos/:id", func(c echo.Context) error {
		id := c.Param("id")
		for _, t := range todos {
			if id == string(t.ID) {
				return c.JSON(http.StatusOK, t)
			}
		}
		return echo.NewHTTPError(http.StatusNotFound, "Todo not found")
	})

	// PUT /todos/:id
	e.PUT("/todos/:id", func(c echo.Context) error {
		id := c.Param("id")
		var updatedTodo todo
		if err := c.Bind(&updatedTodo); err != nil {
			return err
		}
		for i := range todos {
			if id == string(todos[i].ID) {
				todos[i].Title = updatedTodo.Title
				todos[i].Done = updatedTodo.Done
				return c.JSON(http.StatusOK, todos[i])
			}
		}
		return echo.NewHTTPError(http.StatusNotFound, "Todo not found")
	})

	// DELETE /todos/:id
	e.DELETE("/todos/:id", func(c echo.Context) error {
		id := c.Param("id")
		for i := range todos {
			if id == string(todos[i].ID) {
				todos = append(todos[:i], todos[i+1:]...)
				return c.NoContent(http.StatusNoContent)
			}
		}
		return echo.NewHTTPError(http.StatusNotFound, "Todo not found")
	})

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
