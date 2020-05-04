package atcoder

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// IoReadWriter is a utility focusing on mocking stdin/out.
type IoReadWriter struct {
	Reader io.Reader
	Writer io.Writer
}

// SetStringAsReader creates and sets a string reader as input reader.
func (rw *IoReadWriter) SetStringAsReader(str string) {
	rw.Reader = strings.NewReader(str)
}

// ReadStrings parses lines of inputs into 2D string array.
func (rw *IoReadWriter) ReadStrings() [][]string {
	out := [][]string{}
	scanner := bufio.NewScanner(rw.Reader)

	for scanner.Scan() {
		out = append(out, strings.Split(scanner.Text(), " "))
	}

	return out
}

// ReadInts parses lines of inputs into 2D int array.
func (rw *IoReadWriter) ReadInts() [][]int {
	out := [][]int{}
	scanner := bufio.NewScanner(rw.Reader)

	for scanner.Scan() {
		var ints []int

		for _, s := range strings.Split(scanner.Text(), " ") {
			i, err := strconv.Atoi(s)

			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to parse string into int.\n", err)
				os.Exit(1)
			}

			ints = append(ints, i)
		}

		out = append(out, ints)
	}

	return out
}
