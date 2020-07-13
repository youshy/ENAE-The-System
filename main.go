package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	width  int
	height int
	err    error
)

func main() {
	checkTerminalSize()

	done := make(chan bool)
	go play(done)
	go printer(done)
	<-done
}

func checkTerminalSize() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	valuePairs := strings.Fields(string(out))
	height, err = strconv.Atoi(valuePairs[0])
	if err != nil {
		log.Fatal(err)
	}
	width, err = strconv.Atoi(valuePairs[1])
	if err != nil {
		log.Fatal(err)
	}
}
