package config

import (
	"log"

	"github.com/spf13/viper"
)

// TODO: Error handling in general format
func ViperConfigVariable(key string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath("/home/anduser/Documents/RoomBooking/config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file: %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value

}
