package main

import (
	"bufio"
	"os"
	"strconv"
)

type InputReader struct {
	scanner *bufio.Scanner
}

func NewInputReader() InputReader {
	return InputReader{bufio.NewScanner(os.Stdin)}
}

func (ir *InputReader) NextString() (str string, err error) {
	ok := ir.scanner.Scan()
	if !ok {
		err = ir.scanner.Err()
		return
	}

	str = ir.scanner.Text()

	return
}

func (ir *InputReader) NextInt() (num int, err error) {
	str, err := ir.NextString()
	if err != nil {
		return
	}

	num, err = strconv.Atoi(str)

	return
}
