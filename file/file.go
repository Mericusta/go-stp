package stpfile

import (
	"bufio"
	"io"
	"os"
)

// ReadFileLineOneByOne 逐行读取文件内容，执行函数返回 true 则继续读取，返回 false 则结束读取
func ReadFileLineOneByOne(filename string, f func(string) bool) error {
	file, openError := os.Open(filename)
	if openError != nil {
		return openError
	}
	defer file.Close()

	return ReadContentLineOneByOne(file, f)
}

// ReadContentLineOneByOne 逐行读取指定内容，执行函数返回 true 则继续读取，返回 false 则结束读取
func ReadContentLineOneByOne(reader io.Reader, f func(string) bool) error {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if !f(scanner.Text()) {
			break
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}
