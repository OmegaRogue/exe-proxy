package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var silent = false

func main() {

	name := os.Args[0]
	args := os.Args[1:]
	extPos := strings.LastIndex(name,".")
	var shortName string
	var longName string
	if extPos != -1 && extPos != 0 {
		shortName = name[:extPos]
	} else {
		shortName = name
	}


	if strings.HasSuffix(shortName, "_s") {
		silent = true
		shortName = strings.TrimSuffix(shortName, "_s")
	}
	if extPos != -1 && extPos != 0 {
		longName = shortName + name[extPos:]
	} else {
		longName = shortName
	}

	if !silent {
		if len(os.Args) > 1 {
			log.Printf("Running Program \"%v\" with Arguments \"%v\"...\n", longName, args)
		} else {
			log.Printf("Running Program \"%v\"...\n", longName)
		}
	}
	var wg sync.WaitGroup

	c := prepare(append([]string{longName}, args...))
	c.Dir = os.TempDir()
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Run()
	}()
	wg.Wait()
}

func prepare(args []string) *exec.Cmd {
	args = append([]string{"/C"}, args...)
	return exec.Command("cmd", args...)
}