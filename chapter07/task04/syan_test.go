package syan

import (
	"errors"
	"io"
	"testing"
)

var htmlDoc = `
<html>
	<body>
		<h1>Head</h1>
		<a href="ya.ru">яндекс</a>
		<h2>Head2</h2>
		<a href="goo.com">google</a>
	</body>
</html>
`

func TestReader(t *testing.T) {
	tests := []struct {
		err  error
		link string
	}{{nil, "ya.ru"}, {nil, "goo.com"}, {io.EOF, ""}}

	r := NewReader(htmlDoc)
	buf := make([]byte, 512)
	for _, test := range tests {
		n, err := r.Read(buf)
		link := string(buf[:n])
		if !errors.Is(err, test.err) || n != len(test.link) || link != test.link {
			s := "wait: err = %v, n = %d, link = %s\n"
			s += "got: err = %v, n = %d, link = %s\n"
			t.Errorf(s, test.err, len(test.link), test.link, err, n, link)
		}
	}
}

func TestReaderNeg(t *testing.T) {
	tests := []struct {
		err  error
		link string
	}{
		{nil, "ya.ru"},
		{ErrSmallBuffer, ""},
	}

	r := NewReader(htmlDoc)
	buf := make([]byte, 5)
	for _, test := range tests {
		n, err := r.Read(buf)
		link := string(buf[:n])
		if !errors.Is(err, test.err) || n != len(test.link) || link != test.link {
			s := "wait: err = %v, n = %d, link = %s\n"
			s += "got: err = %v, n = %d, link = %s\n"
			t.Errorf(s, test.err, len(test.link), test.link, err, n, link)
		}
	}
}
