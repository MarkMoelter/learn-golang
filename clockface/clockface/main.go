package main

import (
	"os"
	"time"

	"github.com/MarkMoelter/learn-golang/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
