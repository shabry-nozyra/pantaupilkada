package env

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"reflect"
)

type Config struct {
	PostgresHost string `mapstructure:"POSTGRES_HOST" default:""`
	PostgresPort int    `mapstructure:"POSTGRES_PORT" default:"5434"`
	PostgresDB   string `mapstructure:"POSTGRES_DB" default:""`
	PostgresUser string `mapstructure:"POSTGRES_USER" default:""`
	PostgresPass string `mapstructure:"POSTGRES_PASSWORD" default:""`
	AppHost 	 string `mapstructure:"APP_HOST" default:"127.0.0.1:8080"`
}

func (c *Config) ConnectionString() string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		os.Getenv("PostgresUser"),
		os.Getenv("PostgresPass"),
		os.Getenv("PostgresHost"),
		c.PostgresPort,
		os.Getenv("PostgresDB"))
}

var (
	cfg *Config = nil
)

func Get() *Config {
	if cfg == nil {
		cfg = new(Config)

		viper.SetConfigType("yaml")
		viper.SetConfigFile("env.yaml")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()

		_ = viper.ReadInConfig()

		e := reflect.ValueOf(cfg).Elem()
		t := e.Type()
		for i := 0; i < e.NumField(); i++ {
			key := t.Field(i).Tag.Get("mapstructure")
			value := t.Field(i).Tag.Get("default")

			viper.SetDefault(key, value)
		}

		_ = viper.Unmarshal(cfg)
	}

	return cfg
}
