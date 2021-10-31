package main

import (
  _ "github.com/hisamcode/belajar-hexagonal-golang/tzinit"
  "github.com/hisamcode/belajar-hexagonal-golang/app"
  "github.com/hisamcode/belajar-hexagonal-golang/logger"
)

func main() {

  logger.Info("Starting the application")
	app.Start()

}
