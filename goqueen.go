package main

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"github.com/thoas/stats"
	"net/http"
)

type card struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Pin       string `json:"pin"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type log struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	ValidPin  string `json:"valid_pin"`
	CreatedAt string `json:"created_at"`
}

type scheudle struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Mon  int    `json:"mon"`
	Tue  int    `json:"tue"`
	Wed  int    `json:"wed"`
	Thu  int    `json:"thu"`
	Fri  int    `json:"fri"`
	Sat  int    `json:"sat"`
	Sun  int    `json:"sun"`
}

var cards map[string]card

func main() {
	e := echo.New()

	e.Use(mw.Logger)
	e.Use(cors.Default().Handler)

	s := stats.New()
	e.Use(s.Handler)

	e.Get("/stats", func(c *echo.Context) {
		c.JSON(200, s.Data())
	})

	e.Run(":2020")
}
