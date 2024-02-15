package config

import (
	"crud/app/logger"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var configs *viper.Viper

func init() {
	workingDir, _ := os.Getwd()
	resourcesDir := workingDir + "/resources"
	if _, err := os.Stat(resourcesDir + "/app.yaml"); errors.Is(err, os.ErrNotExist) {
		return
	}
	c := loadConfig(resourcesDir, "app.yaml")
	var profiles []string

	if c.IsSet("application.profiles") {
		profiles = strings.Split(get(c, "application.profiles"), ",")
	}
	configs = viper.New()
	addToConfig(configs, c)
	loadAllProfiles(profiles, resourcesDir)
	profiles = append(profiles, "default")
	logger.Info("lodading profiles : " + strings.Join(profiles, ","))
}
func loadConfig(resourcesDir string, configname string) *viper.Viper {
	c := viper.New()
	c.SetConfigName(configname)
	c.SetConfigType("yaml")
	c.AddConfigPath(resourcesDir)
	err := c.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return c
}
func loadAllProfiles(profiles []string, resourcesDir string) {
	if len(profiles) == 0 {
		return
	}
	for _, profile := range profiles {
		otherConfigFileName := "/app-" + profile + ".yaml"
		if _, err := os.Stat(resourcesDir + otherConfigFileName); !errors.Is(err, os.ErrNotExist) {
			addToConfig(configs, loadConfig(resourcesDir, otherConfigFileName))
		}
	}
}

func addToConfig(configs *viper.Viper, c *viper.Viper) {
	keys := c.AllKeys()
	for _, key := range keys {
		configs.Set(key, get(c, key))
	}
}

func get(c *viper.Viper, key string) string {
	val := c.GetString(key)
	if strings.HasPrefix(val, "${") && strings.HasSuffix(val, "}") {
		val, _ = strings.CutPrefix(val, "${")
		val, _ = strings.CutSuffix(val, "}")
		parts := strings.Split(val, ":")
		valFromEnv := os.Getenv(parts[0])
		if valFromEnv == "" && len(parts) == 1 {
			panic("no default value for key " + key + ", no environment variable set for " + parts[0])
		}
		if valFromEnv != "" {
			return valFromEnv
		}
		return parts[1]
	}
	return val
}
func Get(key string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return configs.GetString(key)
}
