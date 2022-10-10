package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	name := os.Args[1]
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")
	totalLines := int64(len(lines))
	conf := strings.Split(lines[0], " ")
	linesPerFrame, err := strconv.ParseInt(conf[0], 10, 16)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	frameDuration, err := strconv.ParseInt(conf[1], 10, 16)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print("\033[H\033[2J")
	s := int64(1)
	totalFrames := (totalLines - 1) / linesPerFrame
	for f := int64(0); f < totalFrames; f++ {
		for fh := (f * linesPerFrame) + s; fh < (f*linesPerFrame+s)+linesPerFrame; fh++ {
			fmt.Fprintf(os.Stdout, "%s\n", strings.ReplaceAll(strings.ReplaceAll(lines[fh], "\\x1b[", "\x1b["), "\\033[", "\033["))
		}
		time.Sleep(time.Duration(frameDuration) * time.Millisecond)
		if f+1 < totalFrames {
			fmt.Print("\033[H\033[2J")
		}
	}

}
