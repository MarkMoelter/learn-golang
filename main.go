package main

import (
	"os"

	"example.com/myproject/countdown"
)

func main() {
	sleeper := &countdown.DefaultSleeper{}
	countdown.Countdown(os.Stdout, sleeper)
}
