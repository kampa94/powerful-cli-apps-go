package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("Word1 Word2 Word3\n")
	exp := 3
	res := count(b, false, false, false)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("Word1\n Word2 Word3\n")
	exp := 2
	res := count(b, true, false, false)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountRunes(t *testing.T) {
	b := bytes.NewBufferString("Word")
	exp := 4
	res := count(b, false, true, false)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("Word")
	exp := 4
	res := count(b, false, false, true)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
