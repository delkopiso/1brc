package calculations

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func CalculateNaive(filename string, w io.Writer) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var stationNames []string
	metrics := map[string]Metric{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		stationName, decimal, found := strings.Cut(text, ";")
		if !found {
			panic("found should have been true")
		}
		measurement, parseErr := strconv.ParseFloat(decimal, 65)
		if parseErr != nil {
			panic(parseErr)
		}
		current, exists := metrics[stationName]
		if !exists {
			stationNames = append(stationNames, stationName)
			metrics[stationName] = Metric{
				Min:   measurement,
				Max:   measurement,
				Sum:   measurement,
				Count: 1,
			}
		} else {
			metrics[stationName] = Metric{
				Min:   min(measurement, current.Min),
				Max:   max(measurement, current.Max),
				Sum:   measurement + current.Sum,
				Count: current.Count + 1,
			}
		}
	}

	slices.Sort(stationNames)
	_, _ = w.Write([]byte("{"))
	for i, name := range stationNames {
		metric := metrics[name]
		mean := metric.Sum / float64(metric.Count)
		format := fmt.Sprintf("%s=%.1f/%.1f/%.1f", name, metric.Min, mean, metric.Max)
		_, _ = w.Write([]byte(format))
		if i < len(stationNames)-1 {
			_, _ = w.Write([]byte(", "))
		}
	}
	_, _ = w.Write([]byte("}"))
}
