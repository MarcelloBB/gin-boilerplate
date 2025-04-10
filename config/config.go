package config

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/ini.v1"
)

var cfg *ini.File
var filePath = "config-file.ini"

func InitializeConfig() {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		emptyCfg := ini.Empty()
		err := emptyCfg.SaveTo(filePath)
		if err != nil {
			fmt.Println("Error creating config file:", err)
		}
	}

	var err error
	cfg, err = ini.Load(filePath)
	if err != nil {
		fmt.Println("Error loading config file:", err)
	}
}

func LoadConfigIni(section, key string, defaultValue interface{}) interface{} {
	if cfg == nil {
		fmt.Println("Config not initialized")
		return defaultValue
	}

	sec, err := cfg.GetSection(section)
	if err != nil {
		sec, _ = cfg.NewSection(section)
	}

	value := sec.Key(key).String()
	if value == "" {
		sec.Key(key).SetValue(fmt.Sprintf("%v", defaultValue))
		cfg.SaveTo(filePath)
		return defaultValue
	}

	switch defaultValue.(type) {
	case bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return defaultValue
		}
		return boolValue
	case int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return intValue
	case string:
		return value
	default:
		return defaultValue
	}
}
