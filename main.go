package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "sampleApp"
	app.Usage = "This app echo input arguments"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello World")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
