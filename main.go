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
	done := make(chan bool)
	go play(done)
	go printer(done)
	<-done
}

func checkTerminalSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return 0, 0
	}
	valuePairs := strings.Fields(string(out))
	height, err = strconv.Atoi(valuePairs[0])
	if err != nil {
		log.Fatal(err)
		return 0, 0
	}
	width, err = strconv.Atoi(valuePairs[1])
	if err != nil {
		log.Fatal(err)
		return 0, 0
	}
	return width, height
}
