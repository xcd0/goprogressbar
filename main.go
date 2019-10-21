package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/nsf/termbox-go"
)

func main() {
	flag.Parse()

	progressVal, _ := strconv.ParseFloat(flag.Arg(0), 64)
	val := progressVal / 1000

	if progressVal > 1000 {
		fmt.Println("Finish!")
		os.Exit(0)
	}

	if progressVal < 0 {
		fmt.Println("Error!")
		os.Exit(1)
	}

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	w, _ := termbox.Size()
	termbox.Close()

	var width int
	width = int(w / 3 * 2)

	var progress int
	progress = int(float64(width) * val)

	f := float64(progress) / float64(width) * 100
	bar := "  ( "
	if f < 100 {
		bar += " "
	}
	bar += strconv.FormatFloat(f, 'f', 1, 64)

	bar += " % ) ["
	for i := 0; i < width; i++ {
		if i < progress {
			bar += "#"
		} else if i == progress {
			bar += ">"
		} else {
			bar += " "
		}
	}
	bar += "]"

	fmt.Printf("\r%v", bar)
}
