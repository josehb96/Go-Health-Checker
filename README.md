# Go Health Checker

![Go](https://img.shields.io/badge/Go-1.20-blue.svg)
![Build](https://img.shields.io/badge/Build-Passing-brightgreen.svg)

Go Health Checker is a simple command-line tool built in Golang to check the health status of websites. Given a domain, it determines whether the website is live or down. 

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Development](#development)

## Features

- Check if a website is up or down by providing the domain name.
- Optionally specify a port to check; defaults to port 80 if not provided.
- Simple and intuitive CLI interface.
- Quick and efficient, leveraging Go's networking capabilities.

## Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/yourusername/go-health-checker.git
   cd go-health-checker
   ```

2. **Build the Application:**

   Ensure you have Go installed. Then, build the application using:

   ```bash
   go build -o healthchecker main.go check.go
   ```

## Usage

The `healthchecker` command requires at least the `--domain` flag. Optionally, you can provide the `--port` flag.

### Command-line Flags

- `--domain, -d` (required): The domain name to check.
- `--port, -p` (optional): The port number to check (defaults to 80).

### Running the Tool

```bash
./healthchecker --domain example.com
```

or

```bash
./healthchecker -d example.com -p 443
```

## Examples

### Check the Status of a Domain

To check if `example.com` is reachable on port 80:

```bash
./healthchecker --domain example.com
```

Output:
```
[UP] example.com is reachable
```

### Check the Status of a Domain on a Specific Port

To check if `example.com` is reachable on port 443:

```bash
./healthchecker -d example.com -p 443
```

Output:
```
[UP] example.com is reachable
```

### When the Website is Down

If a website is down, you might see:

```
[DOWN] example.com is unreachable
```

## Development

### Code Overview

#### `main.go`

Defines the CLI application using the `urfave/cli` package. It handles command-line parsing and invokes the `Check` function to perform the health check.

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2" // Package for creating command-line applications
)

// Main function of the program
func main() {
	// Configuration of the CLI application
	app := &cli.App{
		Name:  "Healthchecker",
		Usage: "A tiny tool that checks whether a website is running or is down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check",
				Required: false,
			},
		},
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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
```

#### `check.go`

Contains the `Check` function, which attempts to establish a TCP connection to the specified domain and port. It returns the status of the connection.

```go
package main

import (
	"fmt"
	"net"
	"time"
)

func Check(domain, port string) string {
	address := domain + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable\nError: %v", domain, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable\nFrom: %v", domain, conn.LocalAddr())
		conn.Close()
	}
	return status
}
```

### Running Tests

To test the application, you can manually run the built executable with various domains and ports to see if the outputs match your expectations.
