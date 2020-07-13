package main

func main() {
	done := make(chan bool)
	go play(done)
	go printer(done)
	<-done
}
