package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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

	for _, element := range slice.elements {
		data, err := os.ReadFile(element)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\t\t%s\n", element, data)
	}
}
