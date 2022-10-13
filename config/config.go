package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MapGeneration mapGenStruct `yaml:"mapGeneration"`
}

type mapGenStruct struct {
	RoomGeneration roomGenStruct `yaml:"roomGeneration"`
}

type roomGenStruct struct {
	MinHeight int `yaml:"minHeight"`
	MaxHeight int `yaml:"maxHeight"`
	MinWidth  int `yaml:"minWidth"`
	MaxWidth  int `yaml:"maxWidth"`
}

var aConfig Config

func init() {
	bytes, err := os.ReadFile("./config/default.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bytes, &aConfig)
	if err != nil {
		panic(err)
	}
}

func Get() Config {
	return aConfig
}

func GetRoomGen() roomGenStruct {
	return aConfig.MapGeneration.RoomGeneration
}
