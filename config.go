package main

import (
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"

	"daqnext/meson-cloud-client/daemon"
)

type AppCfg struct {
	Version  string `yaml:"version"`
	Token    string `yaml:"token"`
	QueryUrl string `yaml:"queryUrl"`
	LogLevel string `yaml:"logLevel"`

	Ipfs daemon.IpfsCfg `yaml:"ipfs"`
}

type AppConfig struct {
	cfg      *AppCfg
	usedFile string
	v        *viper.Viper
}

func loadConfig(extPath string) *AppConfig {
	// Read Config
	v := viper.New()
	v.SetConfigName("config") // name of config file (without extension)
	v.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath(".")      // path to look for the config file in
	v.AddConfigPath(extPath)  // path to look for the config file in

	//viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	if err := v.ReadInConfig(); err != nil {
		log.Panicln("Failed to read the config", err.Error())
	}
	usedFile := v.ConfigFileUsed()

	var C AppCfg
	if err := v.Unmarshal(&C); err != nil {
		log.Panicln("Unable to decode into struct", err.Error())
	}

	return &AppConfig{
		v:        v,
		usedFile: usedFile,
		cfg:      &C,
	}
}

func (c *AppConfig) updateConfig() {

	d, err := yaml.Marshal(c.cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile(c.usedFile, d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
