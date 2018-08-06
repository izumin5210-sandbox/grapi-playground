package main

import (
	"os"

	"github.com/izumin5210-sandbox/grapi-playground/app"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := app.Run()
	if err != nil {
		return 1
	}
	return 0
}
