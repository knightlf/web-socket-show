package conf

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Sever sever
	Web   web
	Log   log
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

var Cfg *Config = nil

func ParserConfig(configPath string) error {
	currentPath, _ := os.Getwd()
	path := currentPath + configPath
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
	//println(Cfg.Sever.Saddr)
	return nil
}
