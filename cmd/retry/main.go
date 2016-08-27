package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/dustin/go-humanize"
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
		/*cli.Float64Flag{
			Name:   "timeout, t",
			Usage:  "maximum seconds per attempt (disabled=0)",
			EnvVar: "RETRY_TIMEOUT",
			Value:  0,
		},*/
		/*cli.Float64Flag{
			Name: "every, e",
			Usage: "ensure is attempt is called every N seconds (similar to cron)",
			EnvVar: "RETRY_EVERY",
			Value: 0,
		},*/
	}

	app.Action = retry
	app.Run(os.Args)
}

func retry(c *cli.Context) error {
	startTime := time.Now()
	attempt := 0

	for {
		attempt++
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
			log.Printf("run %d: command finished with error: %v", attempt, err)
		}
		interval := c.Float64("interval")
		if interval < 0.1 {
			interval = 0.1
		}
		time.Sleep(time.Duration(interval*1000) * time.Millisecond)
	}

	if !c.Bool("quiet") {
		endTime := time.Now()
		totalDuration := humanize.RelTime(endTime, startTime, "", "")
		log.Printf("Command succeeded on attempt %d with a total duration of %s", attempt, totalDuration)
	}

	return nil
}
