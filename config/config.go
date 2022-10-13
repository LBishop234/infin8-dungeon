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
	CircleGeneration circleGenStruct `yaml:"circleGeneration"`
	RectGeneration   rectGenStruct   `yaml:"rectGeneration"`
}

type circleGenStruct struct {
	MinRadius int `yaml:"minRadius"`
	MaxRadius int `yaml:"maxRadius"`
}

type rectGenStruct struct {
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

func GetCircleGenConfig() circleGenStruct {
	return aConfig.MapGeneration.RoomGeneration.CircleGeneration
}

func GetRectGenConfig() rectGenStruct {
	return aConfig.MapGeneration.RoomGeneration.RectGeneration
}
