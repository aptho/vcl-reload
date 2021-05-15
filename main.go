package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	c := make(chan struct{})
	dir := "/etc/varnish"

	go watchDir(dir, c)

	for {
		select {
		case _, ok := <-c:
			err := reload()

			if err != nil {
				fmt.Println(err.Error())
			}

			if !ok {
				fmt.Println("stopped")
			}
		}
	}
}

func reload() error {
	// File has changed, reload varnish
	cmd := exec.Command("varnishreload")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func watchDir(dir string, c chan<- struct{}) {
	lastReload := time.Now()

	for {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			// If this file has a more recent updated time than the last reload time
			if file.ModTime().After(lastReload) {
				fmt.Printf("Changes found in %s...\n", file.Name())

				// Set the last reload time to now
				lastReload = time.Now()

				// Send a notification to reload varnish
				c <- struct{}{}
			}
		}
	}
}
