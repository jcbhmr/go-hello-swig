package main

import (
	"log"
	"log/slog"
	"os"
	"slices"

	"github.com/jcbhmr/go-hello-swig/greet"
)

func init() {
	log.SetFlags(0)
	if slices.Contains([]string{"1", "true"}, os.Getenv("DEBUG")) {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}

func main() {
	slog.Debug("greeting", "name", "Alan Turing")
	greeting := greet.Greeting("Alan Turing")
	log.Printf("greeting: %#+v\n", greeting)

	slog.Debug("greet", "name", "Ada Lovelace")
	greet.Greet("Ada Lovelace")
}
