package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"1brc/pkg/calculations"
)

type Calculator func(string, io.Writer)

var calculators = map[string]Calculator{
	"naive": calculations.CalculateNaive,
}

func main() {
	var use, inFilename, resultFilename, timingFilename string
	flag.StringVar(&use, "use", "naive", "which calculation function to use")
	flag.StringVar(&inFilename, "in", "measurements.txt", "location of input file")
	flag.StringVar(&resultFilename, "result", "result.txt", "location to write result")
	flag.StringVar(&timingFilename, "timing", "timing.txt", "location to write timing")
	flag.Parse()

	calculator, ok := calculators[use]
	if !ok {
		panic(fmt.Sprintf("%q is an unknown function", use))
	}

	resultFile, resultFileErr := os.OpenFile(resultFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if resultFileErr != nil {
		panic(resultFileErr)
	}
	defer func() {
		_ = resultFile.Close()
	}()

	start := time.Now()
	calculator(inFilename, resultFile)
	duration := time.Now().Sub(start)

	timingFile, timingFileErr := os.OpenFile(timingFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if timingFileErr != nil {
		panic(timingFileErr)
	}

	defer func() {
		_ = timingFile.Close()
	}()

	if _, writeErr := timingFile.WriteString(duration.String()); writeErr != nil {
		panic(writeErr)
	}
}
