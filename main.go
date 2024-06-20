package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2" // Package for creating command-line applications
)

func main() {

	// Configuration of the CLI application
	app := &cli.App{
		// Name of the application
		Name: "Healthchecker",
		// Description of the application
		Usage: "A tiny tool that checks whether a website is running or is down",
		// Definition of the CLI flags
		Flags: []cli.Flag{
			// Flag for specifying the domain to check
			&cli.StringFlag{
				Name:     "domain",               // Flag name
				Aliases:  []string{"d"},          // Aliases (alternative names)
				Usage:    "Domain name to check", // Description of the flag
				Required: true,                   // This flag is required
			},
			// Flag for specifying the port to check
			&cli.StringFlag{
				Name:     "port",                 // Flag name
				Aliases:  []string{"p"},          // Aliases (alternative names)
				Usage:    "Port number to check", // Description of the flag
				Required: false,                  // This flag is optional
			},
		},
		// Function executed when the CLI application is called
		Action: func(c *cli.Context) error {

			port := c.String("port")

			if c.String("port") == "" {
				port = "80"
			}

			status := Check(c.String("domain"), port)

			fmt.Println(status)

			return nil
		},
	}

	// Run the CLI application with the provided arguments
	err := app.Run(os.Args)
	// If an error occurs while running the application, log and exit
	if err != nil {
		log.Fatal(err)
	}
}
