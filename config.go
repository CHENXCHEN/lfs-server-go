package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"strings"
)

// Configuration holds application configuration. Values will be pulled from
// environment variables, prefixed by keyPrefix. Default values can be added
// via tags.
type Configuration struct {
	Listen      string `config:"tcp://:8080"`
	Host        string `config:"localhost:8080"`
	ExtOrigin   string `config:""` // consider lfs-test-server may behind a reverse proxy
	MetaDB      string `config:"lfs.db"`
	ContentPath string `config:"lfs-content"`
	AdminUser   string `config:""`
	AdminPass   string `config:""`
	Cert        string `config:""`
	Key         string `config:""`
	Scheme      string `config:"http"`
	Public      string `config:"public"`
	UseTus      string `config:"false"`
	TusHost     string `config:"localhost:1080"`
}

func (c *Configuration) IsHTTPS() bool {
	return strings.Contains(Config.Scheme, "https")
}

func (c *Configuration) IsPublic() bool {
	switch Config.Public {
	case "1", "true", "TRUE":
		return true
	}
	return false
}

func (c *Configuration) IsUsingTus() bool {
	switch Config.UseTus {
	case "1", "true", "TRUE":
		return true
	}
	return false
}

// Config is the global app configuration
var Config = &Configuration{}

func init() {
	configFile := os.Getenv("LFS_SERVER_GO_CONFIG")
	if configFile == "" {
		fmt.Println("LFS_SERVER_GO_CONFIG is not set, Using default config.ini")
		configFile = "config.ini"
	}

	configuration := &Configuration{
		Listen:      "tcp://:8080",
		Host:        "localhost:8080",
		ExtOrigin:   "",
		MetaDB:      "lfs.db",
		ContentPath: "lfs-content",
		AdminUser:   "admin",
		AdminPass:   "admin",
		Cert:        "",
		Key:         "",
		Scheme:      "http",
		Public:      "true",
		UseTus:      "false",
		TusHost:     "localhost:1080",
	}
	cfg, err := ini.Load(configFile)
	if err != nil {
		//panic(fmt.Sprintf("unable to read config from %s, %v", configFile, err))
		fmt.Printf("unable to read config from %s, use default config..., %v\n", configFile, err)
	} else {
		err = cfg.Section("Main").MapTo(configuration)
		if err != nil {
			panic(fmt.Sprintf(fmt.Sprintf("unable to load config.ini[Main], %v", err)))
		}
	}
	Config = configuration
}
