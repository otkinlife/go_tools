package main

import (
	"os"
	"runtime/trace"
)

type Trace struct {
	outFile string
	file    *os.File
}

func New(outFile string) *Trace {
	return &Trace{outFile: outFile}
}

func (t *Trace) Start() error {
	f, err := os.Create(t.outFile)
	if err != nil {
		return err
	}
	t.file = f
	return trace.Start(f)
}

func (t *Trace) Stop() {
	trace.Stop()
	if t.file != nil {
		_ = t.file.Close()
	}
}
