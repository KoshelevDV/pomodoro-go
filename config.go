package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Work       int
	Rest       int
	Long_rest  int
	WorkPeriod int
	WorkCount  int
}

func Init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	config := &Config{
		Work:       25,
		Rest:       5,
		Long_rest:  15,
		WorkPeriod: 4,
		WorkCount:  0,
	}
	file, _ := json.MarshalIndent(config, "", " ")
	exPath = filepath.Join(exPath, "data.json")
	err = ioutil.WriteFile(exPath, file, 0644)
	if err != nil {
		panic(err)
	}
}

func Check() {}

func Load() Config {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	dataPath := filepath.Join(exPath, "data.json")
	data, err := ioutil.ReadFile(dataPath)
	if err != nil {
		panic(err)
	}
	config := Config{}
	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		panic(err)
	}
	return config
}
