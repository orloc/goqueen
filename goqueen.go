package main

import (
	"fmt"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
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
	Assets string `json:"assets"`
}

var cards map[string]card

var configPT *AppConfig = new(AppConfig)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func loadConfig(path string, config *AppConfig) {
	log.Print("Loading configuration...")
	dat, err := ioutil.ReadFile(path)
	checkErr(err)
	fmt.Print(string(dat))

}

func getArgs() string {
	// file
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf("Must specify asset location\n\nUsage: %s [asset_path]\n", os.Args[0])
		os.Exit(1)
	}

	return args[0]
}

func main() {

	configPath := getArgs()

	fmt.Print(configPath)

	loadConfig(configPath, configPT)

	e := echo.New()

	e.Use(mw.Logger)
	e.Use(cors.Default().Handler)

	s := stats.New()
	e.Use(s.Handler)

	e.Get("/", func(c *echo.Context) {
	})

	e.Get("/stats", func(c *echo.Context) {
		c.JSON(200, s.Data())
	})

	e.Run(":8080")
}
