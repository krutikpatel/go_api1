package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port      string
	LogConfig LogConfig
}

type LogConfig struct {
	Level      string `yaml:"level" env:"LOG_LEVEL" default:"info"`
	Format     string `yaml:"format" env:"LOG_FORMAT" default:"json"`
	Output     string `yaml:"output" env:"LOG_OUTPUT" default:"stdout"`
	EnableFile bool   `yaml:"enable_file" env:"LOG_ENABLE_FILE" default:"false"`
	FilePath   string `yaml:"file_path" env:"LOG_FILE_PATH" default:"logs/app.log"`
}

func Load() *Config {
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("LOG_FORMAT", "json")
	viper.SetDefault("LOG_OUTPUT", "stdout")
	viper.SetDefault("LOG_ENABLE_FILE", false)
	viper.SetDefault("LOG_FILE_PATH", "logs/app.log")

	viper.AutomaticEnv()

	port := viper.GetString("PORT")
	logConfig := LogConfig{
		Level:      viper.GetString("LOG_LEVEL"),
		Format:     viper.GetString("LOG_FORMAT"),
		Output:     viper.GetString("LOG_OUTPUT"),
		EnableFile: viper.GetBool("LOG_ENABLE_FILE"),
		FilePath:   viper.GetString("LOG_FILE_PATH"),
	}

	return &Config{
		Port:      port,
		LogConfig: logConfig,
	}
}
