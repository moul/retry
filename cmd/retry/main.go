package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	succeed := false

	for !succeed {
		cmd := exec.Command(os.Args[1], os.Args[2:]...)
		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		err := cmd.Wait()
		if err != nil {
			log.Printf("Command finished with error: %v", err)
		} else {
			succeed = true
		}
		time.Sleep(1 * time.Second)
	}
}
