package main

import (
	"bytes"
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"github.com/thoas/stats"
	"io/ioutil"
	"log"
	"os"
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

type cardlog struct {
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

type AppConfig struct {
	AssetPath string `valid:"required"`
}

func (config AppConfig) getAsset(path string) string {

	var buffer bytes.Buffer

	buffer.WriteString(config.AssetPath)
	buffer.WriteString("/")
	buffer.WriteString(path)

	return buffer.String()
}

var cards map[string]card

func getArgs() string {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf("Must specify asset location\n\nUsage: %s [asset_path]\n", os.Args[0])
		os.Exit(1)
	}

	return args[0]
}

func loadConfig(path string, config *AppConfig) {
	dat, err := ioutil.ReadFile(path)
	checkErr(err)

	jsonErr := json.Unmarshal(dat, config)
	checkErr(jsonErr)

	_, validErr := govalidator.ValidateStruct(*config)
	checkErr(validErr)

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func handleOptions(c *echo.Context) {
}

func main() {
	config := &AppConfig{AssetPath: ""}

	configPath := getArgs()
	loadConfig(configPath, config)

	log.Print("Configuration Loaded!")

	e := echo.New()

	e.Use(mw.Logger)
	e.Use(cors.Default().Handler)

	s := stats.New()
	e.Use(s.Handler)

	e.Index(config.getAsset("index.html"))
	e.Static("/styles", config.AssetPath+"/styles")
	e.Static("/images", config.AssetPath+"/images")
	e.Static("/scripts", config.AssetPath+"/scripts")
	e.Static("/views", config.AssetPath+"/views")

	e.Options("/api/cards", handleOptions)
	e.Options("/api/cards/*", handleOptions)
	e.Options("/api/logs", handleOptions)
	e.Options("/api/schedules", handleOptions)
	e.Options("/api/schedules/*", handleOptions)

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

	e.Get("/stats", func(c *echo.Context) {
		c.JSON(200, s.Data())
	})

	e.Run(":8080")
}
