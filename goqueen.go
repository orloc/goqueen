package main

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	app "github.com/orloc/goqueen/app"
	model "github.com/orloc/goqueen/model"
	"github.com/rs/cors"
	"github.com/thoas/stats"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var cards map[string]model.Card

func handleOptions(c *echo.Context) {
}

func main() {
	config := new(model.AppConfig)

	configPath := app.GetArgs()
	app.LoadConfig(configPath, config)

	log.Print("Configuration Loaded!")

	e := echo.New()

	e.Use(mw.Logger)
	e.Use(cors.Default().Handler)

	s := stats.New()
	e.Use(s.Handler)

	e.Index(config.GetAsset("index.html"))
	e.Static("/styles", config.GetAsset("/styles"))
	e.Static("/images", config.GetAsset("/images"))
	e.Static("/scripts", config.GetAsset("/scripts"))
	e.Static("/views", config.GetAsset("/views"))

	e.Options("/api/cards", handleOptions)
	e.Options("/api/cards/*", handleOptions)
	e.Options("/api/logs", handleOptions)
	e.Options("/api/schedules", handleOptions)
	e.Options("/api/schedules/*", handleOptions)

	// Gets
	e.Get("/api/cards", func(c *echo.Context) {
	})
	e.Get("/api/cards/*", func(c *echo.Context) {
	})
	e.Get("/api/logs", func(c *echo.Context) {
	})
	e.Get("/api/schedules", func(c *echo.Context) {
	})
	e.Get("/api/schedules/*", func(c *echo.Context) {
	})

	// Posts
	e.Post("/api/cards", func(c *echo.Context) {
	})
	e.Post("/api/schedules", func(c *echo.Context) {
		schedule := new(model.Schedule)

		body, err := ioutil.ReadAll(c.Request.Body)
		app.CheckErr(err)

		jsonErr := json.Unmarshal(body, schedule)
		app.CheckErr(jsonErr)

		fmt.Printf("%+v\n", schedule)
		os.Exit(1)

		if err := c.Bind(schedule); err == nil {
			sch := *schedule
			fmt.Printf("%+v\n", sch)

		}

		http.Error(c.Response, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

	})

	e.Get("/stats", func(c *echo.Context) {
		c.JSON(200, s.Data())
	})

	e.Run(":8080")
}
