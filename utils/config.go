package utils

import (
	"encoding/json"
	"log"
	"os"
)

type AppConfigData struct {
	TgBotToken string `json:"TgBotToken"`
	AuthKey    string
	Admins     string
}

var (
	AppConfig     = &AppConfigData{}
	AppConfigPath = "config.json"
)

func init() {
	loadConfig()
}

func loadConfig() {
	if AppConfig.TgBotToken != "" {
		return
	}

	data, err := os.ReadFile(AppConfigPath)
	if err != nil {
		log.Println("加载失败: ", AppConfigPath)
		return
	}

	err = json.Unmarshal(data, AppConfig)
	if err != nil {
		log.Println("解析失败: ", AppConfigPath)
	}
}
