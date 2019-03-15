package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
	"flag"
)

func rand100() (int) {
	return rand.Intn(100)
}

func waitRand(ms int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(ms)))
}

func makeDataSlowly(targetChannel chan string) {
	for {
		waitRand(250)

		n := rand100()
		shipped := false
		waiting := false

		// Use select {} to push onto buffered channel only if it wouldn't overflow
		for {			
			select {
				case targetChannel <- strconv.Itoa(n):
					fmt.Println(n, "put into channel")
					shipped = true
				default:
					if (!waiting) {
						fmt.Println(n, "waiting, channel full")
						waiting = true
					}
			}

			if shipped {
				break
			}
		}
	}
}

func redactEvens(dataChannel chan string) {
	data := <- dataChannel

	s, _ := strconv.Atoi(data)
	if (s % 2 == 0) {
		data = "xx"
	}

	fmt.Println(">> ", data, " <<")
}

func main() {

	flagBuffer := flag.String("buffer", "2", "buffer size")
	flag.Parse()

	bufferSize, _ := strconv.Atoi(*flagBuffer)
	dataz := make(chan string, bufferSize)

	go makeDataSlowly(dataz)

	for {
		waitRand(1500)
		
		redactEvens(dataz)
		redactEvens(dataz)
	}

}