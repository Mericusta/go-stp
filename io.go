package stp

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

func ConsoleInput(ctx context.Context, front, endText string) <-chan string {
	textChan := make(chan string, 8)
	formatFront := front + " "

	go func(_ctx context.Context, _front, _endText string, _textChan chan<- string) {
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
	}(ctx, formatFront, endText, textChan)

	return textChan
}
