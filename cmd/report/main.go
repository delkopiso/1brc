package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type Slice struct {
	elements []string
}

func (s *Slice) String() string {
	if s == nil || len(s.elements) == 0 {
		return "[]"
	}
	var builder strings.Builder
	builder.WriteByte('[')
	builder.WriteString(strings.Join(s.elements, ", "))
	builder.WriteByte(']')
	return builder.String()
}

func (s *Slice) Set(value string) error {
	s.elements = append(s.elements, value)
	return nil
}

func main() {
	var slice Slice
	flag.Var(&slice, "file", "location(s) of timing file(s) to include in report")
	flag.Parse()

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, element := range slice.elements {
		data, err := os.ReadFile(element)
		if err != nil {
			panic(err)
		}

		if _, err = fmt.Fprintf(writer, "%s\t%s\n", element, data); err != nil {
			panic(err)
		}
	}
	if err := writer.Flush(); err != nil {
		panic(err)
	}
}
