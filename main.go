package main

import (
	"os"
	"time"

	"github.com/MarkMoelter/learn-golang/countdown"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{Duration: 1 * time.Second, Sleepy: time.Sleep}
	countdown.Countdown(os.Stdout, sleeper)
}
