package main

import (
	"io"
	"sort"
)

type FieldsGetter interface {
	GetFieldsFromCSV(io.Reader) []string
}

type FieldsSorter interface {
	Sort(fields []string, asc bool) []string
}


type fieldsSorter struct{}

func (f fieldsSorter) Sort(fields []string, asc bool) []string {
	in := make([]string, 0, len(fields))
	for _, b := range fields {
		in = append(in, b)
	}
	sort.Slice(in, func(i, j int)bool {
		if asc {
			return in[i] < in[j]
		}
		return in[j] < in[i]
	})
	return in
}

type CSVWorker interface {
	FieldsCount(io.Reader) int
}

type csvWorker struct {
	fs FieldsSorter
	fg FieldsGetter
}

func (c csvWorker) FieldsCount(r io.Reader) int {
	fields := c.fg.GetFieldsFromCSV(r)
	fields = append(fields, "a")
	return len(c.fs.Sort(fields, true))
}