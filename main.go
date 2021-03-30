package main // import "moul.io/retry"

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "retry"
	author := cli.Author{
		Name:  "Manfred Touron",
		Email: "https://moul.io/retry",
	}
	app.Authors = append(app.Authors, &author)
	app.Version = "0.4.0"
	app.Usage = "retry"

	app.Flags = append(app.Flags, &cli.Float64Flag{
		Name:    "interval, n",
		Usage:   "seconds to wait between attempts",
		Value:   1.0,
		EnvVars: []string{"RETRY_INTERVAL"},
	})

	app.Flags = append(app.Flags, &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "show help",
	})

	app.Flags = append(app.Flags, &cli.BoolFlag{
		Name:    "quiet, q",
		Usage:   "don't print errors",
		EnvVars: []string{"RETRY_QUIET"},
	})

	app.Flags = append(app.Flags, &cli.BoolFlag{
		Name:    "clear, c",
		Usage:   "clear screen between each attempts",
		EnvVars: []string{"RETRY_CLEAR"},
	})

	app.Flags = append(app.Flags, &cli.Float64Flag{
		Name:    "timeout, t",
		Usage:   "maximum seconds per attempt (disabled=0)",
		EnvVars: []string{"RETRY_TIMEOUT"},
		Value:   0,
	})

	app.Flags = append(app.Flags, &cli.IntFlag{
		Name:    "max-attempts, m",
		Usage:   "quit after NUM attempts",
		EnvVars: []string{"RETRY_MAX_ATTEMPTS"},
		Value:   0,
	})

	app.Flags = append(app.Flags, &cli.BoolFlag{
		Name:    "reverse-behavior, r",
		Usage:   "inverse behavior, stop on first fail",
		EnvVars: []string{"RETRY_REVERSE_BEHAVIOR"},
	})

	/*flags = append(flags, &cli.Float64Flag{
		Name:    "every, e",
		Usage:   "ensure is attempt is called every N seconds (similar to cron)",
		EnvVars: []string{"RETRY_EVERY"},
		Value:   0,
	})*/

	app.Action = retry
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Runtime error: %v", err)
	}
}

func retry(c *cli.Context) error {
	if c.Args().Len() < 1 {
		return cli.ShowAppHelp(c)
	}

	startTime := time.Now()
	attempt := 0

	command := c.Args()
	slice := command.Slice()
	if len(slice) == 1 {
		slice = []string{"/bin/sh", "-c", slice[0]}
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
		cmd := exec.CommandContext(ctx, slice[0], slice[1:]...)

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
			os.Exit(1)
		}
	}

	return nil
}
