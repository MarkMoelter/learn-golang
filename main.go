package main

import (
	"os"
	"time"

	"example.com/myproject/countdown"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{Duration: 1 * time.Second, Sleepy: time.Sleep}
	countdown.Countdown(os.Stdout, sleeper)
}
