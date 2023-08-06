package main

import "io"

type LimReader struct {
	R io.Reader
	N int64
}

func (l *LimReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return n, err
}

// LimitReader returns a Reader that reads from r
// but stops with EOF after n bytes.
// The underlying implementation is a *LimitedReader.
func LimitReader(r io.Reader, n int64) io.Reader { return &LimReader{r, n} }
