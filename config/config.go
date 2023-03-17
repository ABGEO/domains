package config

import (
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var config *Config //nolint:gochecknoglobals

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type EmailConfig struct {
	From    string
	To      string
	Subject string
}

type RecaptchaConfig struct {
	SiteKey   string `mapstructure:"site_key"`
	SecretKey string `mapstructure:"secret_key"`
}

type Config struct {
	SMTP struct {
		SMTPConfig `mapstructure:",squash"`
	}
	Email struct {
		EmailConfig `mapstructure:",squash"`
	}
	Recaptcha struct {
		RecaptchaConfig `mapstructure:",squash"`
	}
	GACode string `mapstructure:"ga_code"`
}

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "unable to read config")
	}

	if _, err := os.Stat("./config.local.yaml"); err == nil {
		viper.SetConfigName("config.local")
		viper.AddConfigPath(".")

		if err := viper.MergeInConfig(); err != nil {
			return errors.Wrap(err, "unable to merge config")
		}
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&config); err != nil {
		return errors.Wrap(err, "unable to unmarshal config")
	}

	return nil
}

func GetConfig() *Config {
	return config
}
