package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/ikaros/tir/offset"
	"github.com/ikaros/tir/render"
)

const Version string = "0.2.1"

const usageHead string = `
Usage:

  tmg [options] <image>

Description:

  This tool let you render png and jpg in 256 colors on your terminal

`

var (
	offsetX int
	offsetY int
	pSize   string
)

func printUsage() {
	fmt.Fprint(os.Stderr, usageHead)
	fmt.Println("\nOptions:\n")
	flag.PrintDefaults()
	fmt.Print("\n\n")
	fmt.Printf("Version:\n\n  %s\n", Version)
	fmt.Print("\n")
}

func main() {

	flag.StringVar(&pSize, "psize", "small", "Pixel Size [small, big]")
	flag.IntVar(&offsetX, "offsetX", 0, "Offset in chars from left")
	flag.IntVar(&offsetY, "offsetY", 0, "Offset in lines from top")
	flag.Parse()

	if 1 != len(flag.Args()) {
		printUsage()
		os.Exit(0)
	}

	offset.Init(offsetX, offsetY)

	imgPath := flag.Args()[0]

	reader, err := os.Open(imgPath)
	handleError(err)
	defer reader.Close()

	m, _, err := image.Decode(reader)
	handleError(err)

	r := render.XTerm256{}

	switch pSize {
	case "small":
		fmt.Println(r.Render(m))
	case "big":
		fmt.Println(r.RenderBig(m))
	default:
		fmt.Fprintf(os.Stderr, "Unknown pixel size: %s\n", pSize)
		os.Exit(1)
	}
}

func handleError(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}
