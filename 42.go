package main

import (
	"bytes"
	"io"
	"strings"
)

var (
	// XHeader hmmm...
	XHeader = unlock("Rejva")
	// XValue another hmm...
	XValue = unlock("Uver zr! :)")
)

type xReader struct {
	r io.Reader
}

func (x *xReader) Read(p []byte) (n int, err error) {
	n, err = x.r.Read(p)
	for i := 0; i < len(p); i++ {
		switch {
		case p[i] >= 'A' && p[i] < 'N' || p[i] >= 'a' && p[i] < 'n':
			p[i] += 13
		case p[i] > 'M' && p[i] <= 'Z' || p[i] > 'm' && p[i] <= 'z':
			p[i] -= 13
		}
	}
	return
}

func unlock(s string) string {
	r := strings.NewReader(s)
	x := &xReader{r}
	b := &bytes.Buffer{}
	if _, err := b.ReadFrom(x); err != nil {
		return ""
	}
	return b.String()
}
