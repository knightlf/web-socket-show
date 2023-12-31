package conf

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Sever    sever
	Web      web
	Log      log
	Ckconfig ckconfig
}

type sever struct {
	Saddr string `json:"svr_addr"`
	Sport string `json:"svr_port"`
}

type log struct {
	Path  string `json:"log_path"`
	Level string `json:"log_level"`
}

type web struct {
	Waddr string `json:"web_addr"`
	Wport string `json:"web_port"`
}

type ckconfig struct {
	Host             []string `json:"host"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	DatabaseName     string   `json:"databaseName"`
	CompressionLevel int      `json:"compressionLevel"` // 10
	OnCluster        string   `json:"onCluster"`
	UserMaxMemUsage  int64    `json:"userMaxMemUsage"` // 1073741824 = 1G
}

var Cfg *Config = nil

func ParserConfig(configPath string) error {
	currentPath, _ := os.Getwd()
	path := currentPath + configPath
	//println(path)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	if err = decoder.Decode(&Cfg); err != nil {
		fmt.Println("init config error:", err.Error())
		return err
	}
	//println(Cfg.Ckconfig.Host)
	return nil
}
