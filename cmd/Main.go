package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

var (
	spashScreen = `
	🇸​​​​​🇪​​​​​🇷​​​​​🇻​​​​​🇪​​​​​🇷​​​​​ 🇺​​​​​🇵​​​​​
	
	golang seed-go
	https://github.com/sapiderman/mock-server/blob/master/README.md
	`
)

func init() {
	fmt.Println(spashScreen)
	log.Info("intializing")
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

// Main entry point
func main() {

	// start server

}
