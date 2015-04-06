package model

import (
	"bytes"
)

type AppConfig struct {
	AssetPath string `valid:"required"`
	DbName    string `valid:"required"`
}

func (config AppConfig) GetAsset(path string) string {

	var buffer bytes.Buffer

	buffer.WriteString(config.AssetPath)
	buffer.WriteString("/")
	buffer.WriteString(path)

	return buffer.String()
}
