package main

import (
	"flag"
	"fmt"
)

//Args input arguments for Ray Tracer
type Args struct {
	OutDir    string
	Filename  string
	InputFile string
}

func main() {
	fmt.Println("Ray Tracer in a weekend")
	args := defineFlags()

	fmt.Println("Creating Output Directory ", args.OutDir)
	//createOutputDirectory(args.OutDir)

	outputImage(args.Filename)

	flag.Parse()
}

func defineFlags() Args {
	var a = Args{}
	///home/gabe/Documents/census/2016_GCP_ALL_for_AUS_short-header/2016 Census GCP All Geographies for AUST/STE/AUST/2016Census_G02_AUS_STE.csv
	a.OutDir = *flag.String("o", "./out", "output directory")
	a.Filename = *flag.String("f", "image.png", "name of output image in png")
	a.InputFile = *flag.String("i", "", "input filefor any config")

	return a
}
