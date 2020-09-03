package main

import (
	"fmt"
	"log"

	"github.com/abmid/canvas-config-checker/internal/message"
	"github.com/abmid/canvas-config-checker/pkg/apache"
	"github.com/abmid/canvas-config-checker/pkg/canvas"
	"github.com/spf13/viper"
)

func main() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.SetConfigName("settings") // name of config file (without extension)
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	message.Banner()
	canvasNotEquals, canvasGroupErrors, err := canvas.Run(viper)
	if err != nil {
		log.Fatalf(err.Error())
	}

	apacheNotEquals, apacheGroupErrors, err := apache.Run(viper)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if len(canvasNotEquals) == 0 && len(canvasGroupErrors) == 0 && len(apacheNotEquals) == 0 && len(apacheGroupErrors) == 0 {
		message.Ready("production")
	}

	message.SummaryGroupError(canvasGroupErrors, apacheGroupErrors)
	message.SummaryNotEqual(canvasNotEquals, apacheNotEquals)

}
