package counter

import (
	"bufio"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		n++
	}
	*c += WordCounter(n)
	return
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		n++
	}
	*c += LineCounter(n)
	return
}
