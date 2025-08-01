package main

import (
	"os"
	"time"

	"example.com/myproject/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
