package stp

import (
	"bufio"
	"fmt"
	"os"
)

func ConsoleInput(front, endText string) <-chan string {
	textChan := make(chan string, 8)

	go func(_front, _endText string, _textChan chan<- string) {
		fmt.Print(_front)
		var (
			ok  bool
			err error
			in  *bufio.Scanner = bufio.NewScanner(os.Stdin)
		)
		for ok, err = in.Scan(), in.Err(); ok && err == nil; ok, err = in.Scan(), in.Err() {
			text := in.Text()
			if text == _endText {
				break
			}
			_textChan <- text
			fmt.Print(_front)
		}
	}(front, endText, textChan)

	return textChan
}
