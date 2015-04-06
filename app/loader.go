package app

import (
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	model "github.com/orloc/goqueen/model"
	"io/ioutil"
	"os"
)

func GetArgs() string {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf("Must specify asset location\n\nUsage: %s [asset_path]\n", os.Args[0])
		os.Exit(1)
	}

	return args[0]
}

func LoadConfig(path string, config *model.AppConfig) {
	dat, err := ioutil.ReadFile(path)
	CheckErr(err)

	jsonErr := json.Unmarshal(dat, config)
	CheckErr(jsonErr)

	_, validErr := govalidator.ValidateStruct(*config)
	CheckErr(validErr)

}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}
