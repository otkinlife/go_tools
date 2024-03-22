package main

import (
	"log"
	"testing"
)

func TestTrace(t *testing.T) {
	trace := New("trace.out")
	if err := trace.Start(); err != nil {
		t.Fatal(err)
	}
	defer trace.Stop()
	for i := 0; i < 1000; i++ {
		log.Println(i)
	}
}
