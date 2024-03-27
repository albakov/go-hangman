package main

import (
	"github.com/albakov/go-hangman/internal/app"
	"github.com/albakov/go-hangman/internal/config"
)

func main() {
	app.New(config.MustNew()).Start()
}
