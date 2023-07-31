package configuration

import (
	"fmt"
	"github.com/sw90lee/batch-sample/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Logger   Logger
	DBConfig DBConfig
}

type Logger struct {
	Filename string `yaml:"filename"`
	Level    string `yaml:"level"`
	Env      string `yaml:"env"`
}

type DBConfig struct {
	DB_URL      string `yaml:"db_url"`
	DB_USER     string `yaml:"db_user"`
	DB_PASSWORD string `yaml:"db_password"`
	DB_NAME     string `yaml:"db_name"`
	DB_PORT     int    `yaml:"db_port"`
	DB_TABLE    string `yaml:"db_table"`
}

var log = logger.NewLogger()

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}

	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func NewDBConfig() *DBConfig {
	cfg, err := NewConfig("config.yml")
	if err != nil {
		log.Error(err.Error())
	}

	return &DBConfig{
		// getEnv(환경변수 Key값, defaultValue)
		DB_URL:      getEnv("DB_URL", cfg.DBConfig.DB_URL),
		DB_USER:     getEnv("DB_USER", cfg.DBConfig.DB_USER),
		DB_PASSWORD: getEnv("DB_PASSWORD", cfg.DBConfig.DB_PASSWORD),
		DB_NAME:     getEnv("DB_NAME", cfg.DBConfig.DB_NAME),
		DB_PORT:     getEnvAsInt("DB_PORT", cfg.DBConfig.DB_PORT),
		DB_TABLE:    getEnv("test", cfg.DBConfig.DB_TABLE),
	}
}

// env String 반환
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// env int 반환
func getEnvAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	var intValue int
	_, err := fmt.Scan(value, &intValue)
	if err != nil {
		return defaultValue
	}
	return intValue
}
