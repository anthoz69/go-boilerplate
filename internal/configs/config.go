package configs

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func NewViperConfig() *viper.Viper {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		log.Panic(err)
	}

	return v
}
