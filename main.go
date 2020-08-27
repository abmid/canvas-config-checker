package main

import (
	"fmt"
	"log"

	"github.com/abmid/canvas-env-checker/internal/message"
	"github.com/abmid/canvas-env-checker/pkg/canvas"
	"github.com/spf13/viper"
)

type NotEqual struct {
	Group string
}

func Run(viper *viper.Viper) {
	message.Banner()
	notEqual := []NotEqual{}
	canvas := canvas.New(viper)

	isEqual, err := canvas.RunCanvas()
	if err != nil {
		log.Fatalf(err.Error())
	}
	if len(isEqual) > 0 {
		notEqual = append(notEqual, NotEqual{Group: "canvas"})
	}

	if len(notEqual) > 0 {
		log.Fatalf("message")
		// message not ready
	} else {
		message.Ready("Production")
	}
}

func main() {
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.SetConfigName("settings") // name of config file (without extension)
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Run(viper)
}
