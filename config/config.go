package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nekomeowww/vig/logger"
	"gopkg.in/yaml.v2"
)

var (
	// BackendVersion version of backend
	BackendVersion = "1.0.0"

	// FrontendVersion version of frontend
	FrontendVersion = "1.0.0"

	// LastCommit the last commit hash
	LastCommit = ""

	// Stage development stage
	Stage = "test"
)

// Conf 配置
var Conf *Config

// Config vig config
type Config struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

// Init 初始化
func Init() {
	content, err := ioutil.ReadFile(getFilePath())
	if err != nil {
		logger.Error("config file not found, host configuration defaults to http://127.0.0.1:2100")
		Conf = &Config{
			IP:   "127.0.0.1",
			Port: "2100",
		}
	}

	fmt.Println(string(content))
	err = yaml.Unmarshal(content, &Conf)
	if err != nil {
		logger.Fatal(err)
	}
}

func getFilePath() string {
	filePath := "./config.yaml"
	stat, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			filePath = "./config.yml"
		} else {
			logger.Fatal(err)
		}
	}
	stat, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return ""
		}
		logger.Fatal(err)
	}
	if !stat.IsDir() {
		return filePath
	}

	return ""
}
