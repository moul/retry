package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "retry"
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul/retry"
	app.Version = "0.1.0"
	app.Usage = "retry"

	app.Flags = []cli.Flag{
		cli.Float64Flag{
			Name:   "interval, n",
			Usage:  "seconds to wait between attempts",
			Value:  1.0,
			EnvVar: "RETRY_INTERVAL",
		},
		cli.BoolFlag{
			Name:   "quiet, q",
			Usage:  "don't print errors",
			EnvVar: "RETRY_QUIET",
		},
	}

	app.Action = retry
	app.Run(os.Args)
}

func retry(c *cli.Context) error {
	for {
		cmd := exec.Command(c.Args()[0], c.Args()[1:]...)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		err := cmd.Wait()
		if err == nil {
			break
		}

		if !c.Bool("quiet") {
			log.Printf("Command finished with error: %v", err)
		}
		time.Sleep(time.Duration(c.Float64("interval")*1000) * time.Millisecond)
	}
	return nil
}
