package main

import (
	"io"
	"reflect"
	"testing"
)


func TestFieldsSorter( t*testing.T) {
	var fs FieldsSorter
	fieldsIn1 := []string{"b", "c", "a"}
	asc1 := true
	out1 :=[]string{"a","b", "c"}
	fieldsIn2 := []string{"a", "c", "b"}
	asc2 :=  false
	out2 := []string{"c", "b", "a"}
	fs = fieldsSorter{}
	out := fs.Sort(fieldsIn1, asc1)
	if !reflect.DeepEqual(out, out1) {
		t.Errorf("wrong out")
	}
	if reflect.DeepEqual(out, fieldsIn1) {
		t.Errorf("source changed")
	}
	out = fs.Sort(fieldsIn2, asc2)
	if !reflect.DeepEqual(out, out2) {
		t.Errorf("wrong out")
	}
	if reflect.DeepEqual(out, fieldsIn2) {
		t.Errorf("source changed")
	}
}

type stubFieldsGetter struct{}

func (s stubFieldsGetter) GetFieldsFromCSV(reader io.Reader) []string {
	return []string{"c","b", "a"}
}

type mockFieldsSorter struct{
	t *testing.T
}

func newMockFieldsSorter(t *testing.T) *mockFieldsSorter {
	return &mockFieldsSorter{t: t}
}


func (m mockFieldsSorter) Sort(fields []string, asc bool) []string {
	want := stubFieldsGetter{}.GetFieldsFromCSV(nil)
	if !reflect.DeepEqual(want, fields)	{
		m.t.Errorf("wrong argument fields")
	}
	if !asc {
		m.t.Error("wrong argument asc")
	}
	return []string{"a","b","c"}
}

func TestCSVWorker(t *testing.T) {
	w := csvWorker{
		fs: newMockFieldsSorter(t),
		fg: stubFieldsGetter{},
	}
	if w.FieldsCount(nil) != 3 {
		t.Error("wrong count")
	}
}