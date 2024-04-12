package main

import (
	"log"
	"log/slog"
	"os"
	"slices"

	greet "github.com/jcbhmr/go-hello-swig/src"
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
	slog.Info("greeting", "result", greeting)

	slog.Debug("greet", "name", "Ada Lovelace")
	greet.Greet("Ada Lovelace")
}
