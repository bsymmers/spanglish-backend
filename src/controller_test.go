package main

import "testing"

func TestGetData(t *testing.T) {
	got := getData("EN-SP")

	if val, ok := got["binary"]; !ok {
		t.Errorf("Incorrect Map Retrieved, got: %q", val)
	}
}

func TestWordProcessor(t *testing.T) {
	got := wordProcessor("Hello world, I'm Brandon")

	if len(got) != 7 {
		t.Errorf("Word's Processed Incorrectly, got: %q", got)
	}

}
