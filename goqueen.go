package main

import (
	_ "fmt"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	app "github.com/orloc/goqueen/app"
	model "github.com/orloc/goqueen/model"
	"github.com/rs/cors"
	"github.com/thoas/stats"
	"log"
	"net/http"
)

func handleOptions(c *echo.Context) {
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	config := new(model.AppConfig)
	configPath := app.GetArgs()
	app.LoadConfig(configPath, config)

	scheduleManager := app.ScheduleManager{
		DbName: config.DbName, TableName: "schedules",
	}

	scheduleManager.SetupDB(false)

	log.Print("Configuration Loaded!")

	e := echo.New()

	e.Use(mw.Logger)
	e.Use(cors.Default().Handler)

	s := stats.New()
	e.Use(s.Handler)

	// ======= Routes ======= //

	/*
	 * UI Assets
	 */
	e.Index(config.GetAsset("index.html"))
	e.Static("/styles", config.GetAsset("/styles"))
	e.Static("/images", config.GetAsset("/images"))
	e.Static("/scripts", config.GetAsset("/scripts"))
	e.Static("/views", config.GetAsset("/views"))

	/*
	 * Cards
	 */
	e.Get("/api/cards", func(c *echo.Context) {
	})
	e.Get("/api/cards/:id", func(c *echo.Context) {
	})

	e.Post("/api/cards", func(c *echo.Context) {
	})

	e.Options("/api/cards", handleOptions)
	e.Options("/api/cards/*", handleOptions)

	/*
	 * Scheudles
	 */
	e.Get("/api/schedules", func(c *echo.Context) {
		result := scheduleManager.GetAll()
		c.JSON(200, result)
	})

	e.Get("/api/schedules/:id", func(c *echo.Context) {
		schedule := scheduleManager.GetById(c.P(0))
		if schedule.Id == 0 {
			http.Error(c.Response, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			c.JSON(200, schedule)
		}
	})

	e.Post("/api/schedules", func(c *echo.Context) {
		schedule := new(model.Schedule)

		if err := c.Bind(schedule); err == nil {
			sch := *schedule

			scheduleManager.Save(schedule)

			c.JSON(200, sch)
		} else {
			http.Error(c.Response, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
	})
	e.Options("/api/schedules", handleOptions)
	e.Options("/api/schedules/*", handleOptions)

	/*
	 * Logs
	 */
	e.Get("/api/logs", func(c *echo.Context) {
	})
	e.Options("/api/logs", handleOptions)

	/*
	 * Misc
	 */
	e.Get("/stats", func(c *echo.Context) {
		c.JSON(200, s.Data())
	})

	e.Run(":8080")
}
