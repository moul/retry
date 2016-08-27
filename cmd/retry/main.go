package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	for {
		cmd := exec.Command(os.Args[1], os.Args[2:]...)

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

		log.Printf("Command finished with error: %v", err)
		time.Sleep(1 * time.Second)
	}
}
