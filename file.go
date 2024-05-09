package stp

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func writeFileByFileFlag(path string, handler func([]byte) ([]byte, error), flag int) (*os.File, []byte, error) {
	if handler == nil {
		return nil, nil, fmt.Errorf("handler is nil")
	}
	var (
		f       *os.File
		content []byte
		err     error
	)
	if IsExist(path) {
		f, err = os.OpenFile(path, flag, 0666)
	} else {
		f, err = os.Create(path)
	}
	if err != nil {
		return nil, nil, err
	}
	content, err = io.ReadAll(f)
	if err != nil {
		return nil, nil, err
	}
	newContent, err := handler(content)
	if err != nil {
		return nil, nil, err
	}
	return f, newContent, nil
}

// WriteFileByOverwriting 通过覆盖写入文件
func WriteFileByOverwriting(path string, handler func([]byte) ([]byte, error)) error {
	f, newContent, err := writeFileByFileFlag(path, handler, os.O_RDWR|os.O_CREATE)
	if err != nil {
		return err
	}
	writeLen, err := f.WriteAt(newContent, 0)
	if err != nil {
		return err
	}
	return f.Truncate(int64(writeLen))
}

// WriteFileByAppend 通过追加写入文件
func WriteFileByAppend(path string, handler func([]byte) ([]byte, error)) error {
	f, newContent, err := writeFileByFileFlag(path, handler, os.O_RDWR|os.O_CREATE|os.O_APPEND)
	if err != nil {
		return err
	}
	_, err = f.Write(newContent)
	if err != nil {
		return err
	}
	return nil
}

// ReadFileLineOneByOne 逐行读取文件内容，执行函数返回 true 则继续读取，返回 false 则结束读取
func ReadFileLineOneByOne(filename string, f func(string, int) bool) error {
	file, openError := os.Open(filename)
	if openError != nil {
		return openError
	}
	defer file.Close()

	return ReadContentLineOneByOne(file, f)
}

// ReadContentLineOneByOne 逐行读取指定内容，执行函数返回 true 则继续读取，返回 false 则结束读取
func ReadContentLineOneByOne(reader io.Reader, f func(string, int) bool) error {
	index, scanner := 0, bufio.NewScanner(reader)

	for scanner.Scan() {
		if !f(scanner.Text(), index) {
			break
		}
		index++
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}

// IsExist 检查文件或文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// TraverseDirectorySpecificFileWithFunction 遍历文件夹获取所有绑定类型的文件
func TraverseDirectorySpecificFileWithFunction(directory, syntax string, operate func(string, fs.DirEntry) error) error {
	syntaxExt := fmt.Sprintf(".%v", syntax)
	return filepath.WalkDir(directory, func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filePath != directory {
			if d.IsDir() {
				return err
			}
			if path.Ext(filePath) == syntaxExt {
				err := operate(filePath, d)
				return err
			}
		}
		return nil
	})
}

// CreateDir 创建文件夹
func CreateDir(directoryPath string) error {
	err := os.Mkdir(directoryPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// CreateFile 创建文件
func CreateFile(filePath string) (*os.File, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// FormatFilePathWithOS 根据操作系统格式化路径
func FormatFilePathWithOS(filePath string) string {
	osLinux := "linux"
	operationSystem := runtime.GOOS
	beReplaced := "/"
	toReplace := "\\"
	if operationSystem == osLinux {
		beReplaced, toReplace = toReplace, beReplaced
	}
	return strings.ReplaceAll(filePath, beReplaced, toReplace)
}
