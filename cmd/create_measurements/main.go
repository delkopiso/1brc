package main

import (
	"bufio"
	"flag"
	"math"
	"math/rand/v2"
	"os"
	"strconv"
)

func main() {
	var filename string
	var size uint
	flag.StringVar(&filename, "file", "measurements.txt", "location to write records")
	flag.UintVar(&size, "size", 1_000_000_000, "number of records to create")
	flag.Parse()

	if size == 0 {
		panic("size must be a positive integer")
	}

	file, fileErr := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if fileErr != nil {
		panic(fileErr)
	}
	defer func() {
		_ = file.Close()
	}()

	numStations := len(stations)
	writer := bufio.NewWriter(file)
	for range size {
		var err error
		station := stations[rand.IntN(numStations)]
		measurement := rand.NormFloat64()*10 + station.Temp
		rounded := strconv.FormatFloat(math.Round(measurement*10.0)/10.0, 'f', -1, 64)
		_, err = writer.WriteString(station.Id + ";" + rounded)
		if err != nil {
			panic(err)
		}
		_, err = writer.WriteRune('\n')
		if err != nil {
			panic(err)
		}
	}
	if flushErr := writer.Flush(); flushErr != nil {
		panic(flushErr)
	}
}
