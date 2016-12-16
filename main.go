package main

import (
	"os"
	log "github.com/Sirupsen/logrus"
	"github.com/devender/zipcal/loader"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	os.Exit(realMain())
}

func realMain() int {
	sb := loader.Starbucks{Source: "/Users/dgollapally/Downloads/data/starbucks.csv"}

	locations := sb.Load()

	for _,loc := range locations {
		log.Debug(loc)
	}
	return 0
}

