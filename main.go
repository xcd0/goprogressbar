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
	if progressVal < 0 {
		fmt.Println("Error!")
		os.Exit(1)
	} else if progressVal > 1000 {
		progressVal = 1000
	}

	// ウィンドウサイズ取得
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	w, _ := termbox.Size()
	termbox.Close()

	// 進捗バーは画面の半分の長さとする
	width := int(w / 2)

	f := progressVal / 1000
	printnum := int(float64(width) * f)
	fp := f * 100

	// 出力
	bar := "  ( "
	if progressVal < 100 {
		bar += "  "
	} else if progressVal < 1000 {
		bar += " "
	}
	bar += strconv.FormatFloat(fp, 'f', 1, 64)
	bar += " % ) ["

	for i := 0; i < width; i++ {
		if i < printnum {
			bar += "#"
		} else if i == printnum {
			bar += ">"
		} else {
			bar += " "
		}
	}
	bar += "]"

	fmt.Printf("\r%v", bar)
}
