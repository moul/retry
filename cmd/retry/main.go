package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	humanize "github.com/dustin/go-humanize"
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
		cli.BoolFlag{
			Name:   "clear, c",
			Usage:  "clear screen between each attempts",
			EnvVar: "RETRY_CLEAR",
		},
		cli.Float64Flag{
			Name:   "timeout, t",
			Usage:  "maximum seconds per attempt (disabled=0)",
			EnvVar: "RETRY_TIMEOUT",
			Value:  0,
		},
		cli.IntFlag{
			Name:   "max-attempts, m",
			Usage:  "quit after NUM attempts",
			EnvVar: "RETRY_MAX_ATTEMPTS",
			Value:  0,
		},
		cli.BoolFlag{
			Name:   "reverse-behavior, r",
			Usage:  "inverse behavior, stop on first fail",
			EnvVar: "RETRY_REVERSE_BEHAVIOR",
		},
		/*cli.Float64Flag{
			Name: "every, e",
			Usage: "ensure is attempt is called every N seconds (similar to cron)",
			EnvVar: "RETRY_EVERY",
			Value: 0,
		},*/
	}

	app.Action = retry
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Runtime error: %v", err)
	}
}

func retry(c *cli.Context) error {
	if len(c.Args()) < 1 {
		return cli.ShowAppHelp(c)
	}

	startTime := time.Now()
	attempt := 0

	command := c.Args()
	if len(command) == 1 {
		command = []string{"/bin/sh", "-c", command[0]}
	}

	succeed := false
	maxAttempts := c.Int("max-attempts")
	for {
		attempt++
		var ctx context.Context
		if c.Float64("timeout") > 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(context.Background(), time.Duration(c.Float64("timeout"))*time.Second)
			defer cancel()
		} else {
			ctx = context.Background()
		}
		cmd := exec.CommandContext(ctx, command[0], command[1:]...)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		err := cmd.Wait()
		if c.Bool("reverse-behavior") {
			if err != nil {
				succeed = true
				break
			}
		} else {
			if err == nil {
				succeed = true
				break
			}
		}

		if ctx.Err() == context.DeadlineExceeded {
			if !c.Bool("quiet") {
				log.Printf("run %d: command timed out", attempt)
			}
		} else {
			if !c.Bool("quiet") {
				log.Printf("run %d: command finished with error: %v", attempt, err)
			}
		}

		if maxAttempts > 0 && attempt >= maxAttempts {
			break
		}

		interval := c.Float64("interval")
		if interval < 0.1 {
			interval = 0.1
		}
		time.Sleep(time.Duration(interval*1000) * time.Millisecond)
		if c.Bool("clear") {
			fmt.Print("\x1bc")
		}
	}

	if !c.Bool("quiet") {
		endTime := time.Now()
		totalDuration := humanize.RelTime(endTime, startTime, "", "")
		if totalDuration == "now" {
			totalDuration = "0 second"
		}
		fmt.Fprintln(os.Stderr)
		if succeed {
			log.Printf("Command succeeded on attempt %d with a total duration of %s", attempt, totalDuration)
		} else {
			log.Printf("Command failed %d times with a total duration of %s", attempt, totalDuration)
		}
	}

	return nil
}
