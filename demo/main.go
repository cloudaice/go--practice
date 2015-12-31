package main

import (
	"encoding/binary"
	"io"
)

type Gopher struct {
	Name     string
	AgeYears int
}

func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
	err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
	if err != nil {
		return
	}
	size += 4
	n, err := w.Write([]byte(g.Name))
	size += int64(n)
	if err != nil {
		return
	}
	err = binary.Write(w, binary.LittleEndian, int64(len(g.AgeYears)))
	if err == nil {
		size += 4
	}
	return
}

type binWriter struct {
	w    io.Writer
	size int64
	err  error
}

func (w *binWriter) Write(v interface{}) {
	if w.err != nil {
		return
	}
	if w.err = binary.Write(w, binary.LittleEndian, v); w.err == nil {
		w.size += int64(binary.Size(v))
	}
}

func (g *Gopher) WriteTo2(w io.Writer) (int64, error) {
	bw := &binWriter{w: w}
	bw.Write(g.Name)
	bw.Write(g.AgeYears)
	return bw.size, bw.err
}
