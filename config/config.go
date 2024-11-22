package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
	API      APIConfig      `mapstructure:"api"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

var cfg config

var LogDebugf func(format string, v ...interface{})

// Load creates a single
func Load(display bool, debug, debugSQL bool) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	// Incase test cases require loading configs
	viper.AddConfigPath("../config")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.SetConfigType("yaml")

	path := getConfigFile()
	viper.AddConfigPath(path)

	if err := load(); err != nil {
		panic(err)
	}

	if debug {
		SetMode(DebugMode)
	}

}

/* ------------------------------
         Utility Functions
------------------------------ */

func load() error {
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	newCfg := config{}
	err = viper.Unmarshal(&newCfg)
	if err != nil {
		return err
	}

	cfg = newCfg

	return nil
}

func getConfigFile() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	p := filepath.Dir(ex)
	configFilePath := ""
	configFileLocation := "config.yaml"
	for i := 0; i <= 10; i++ {
		configFilePath = filepath.Join(p, configFileLocation)
		if _, err := os.Stat(configFilePath); err == nil {
			break
		}
		p = filepath.Dir(p)
	}
	return p //flag.String("f", configFilePath, "the config file")
}

func Hello(name string) string {
	return "Hello, " + name
}
