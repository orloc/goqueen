package main

import (
	"fmt"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
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

	config := new(app.AppConfig)
	configPath := app.GetArgs()
	app.LoadConfig(configPath, config)

	fmt.Printf("%+v\n", config)

	scheduleManager := app.ScheduleManager{
		DbName:    config.DbName,
		TableName: "schedules",
		Options:   config.DbConfig,
	}

	cardManager := app.CardManager{
		DbName:    config.DbName,
		TableName: "cards",
		Options:   config.DbConfig,
	}

	managers := [...]app.ModelManager{
		scheduleManager,
		cardManager,
	}

	for _, m := range managers {
		m.SetupDB(false)
	}

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
		result := cardManager.GetAll()
		c.JSON(200, result)
	})
	e.Get("/api/cards/:id", func(c *echo.Context) {
		response := scheduleManager.GetById(c.P(0))
		card := response.Card

		if card.Id == 0 {
			http.Error(c.Response, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			c.JSON(200, card)
		}
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
		response := scheduleManager.GetById(c.P(0))
		schedule := response.Schedule

		if schedule.Id == 0 {
			http.Error(c.Response, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			c.JSON(200, schedule)
		}
	})

	e.Put("/api/schedules/:id", func(c *echo.Context) {
		response := scheduleManager.GetById(c.P(0))

		scheduleManager.Update(response, response.Schedule.Id)

		c.JSON(200, response.Schedule)

	})

	e.Post("/api/schedules", func(c *echo.Context) {
		schedule := new(model.Schedule)

		if err := c.Bind(schedule); err == nil {
			sch := *schedule

			request := &app.ModelMap{Schedule: &sch}

			sch.Id = scheduleManager.Save(request)

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

	e.Run(fmt.Sprintf(":%s", config.DaemonPort))
}
