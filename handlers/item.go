package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Define a struct to represent our data model
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Create a slice of Item to store our data
var items []Item

// Handler functions for our routes

func GetItems(c echo.Context) error {
	return c.JSON(http.StatusOK, items)
}

func CreateItem(c echo.Context) error {
	item := new(Item)
	if err := c.Bind(item); err != nil {
		return err
	}
	items = append(items, *item)
	return c.JSON(http.StatusCreated, item)
}

func GetItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for _, item := range items {
		if item.ID == id {
			return c.JSON(http.StatusOK, item)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, "Item not found")
}

func UpdateItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for i, item := range items {
		if item.ID == id {
			newItem := new(Item)
			if err := c.Bind(newItem); err != nil {
				return err
			}
			newItem.ID = id
			items[i] = *newItem
			return c.JSON(http.StatusOK, newItem)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, "Item not found")
}

func DeleteItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return echo.NewHTTPError(http.StatusNotFound, "Item not found")
}