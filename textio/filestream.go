package textio

import (
	"os"
	"bufio"
)

type StringStream struct {
	file *os.File
	scanner *bufio.Scanner
	Success bool
	err error
}

func NewStringStream(filename string) StringStream {
	var ss StringStream

	ss.file, ss.err = os.Open(filename)
	ss.scanner = bufio.NewScanner(ss.file)
	ss.Success = ss.err == nil

	return ss
}

func (this *StringStream) Close() {
	this.file.Close()
}

func (this *StringStream) ReadString(str *string) bool {
	next := this.scanner.Scan()

	if next {
		*str = this.scanner.Text()
		this.Success = true
	} else {
		this.err = this.scanner.Err()
		this.Success = this.err == nil
	}

	return next
}

func (this *StringStream) IsSuccess() bool {
	return this.Success
}

func (this *StringStream) GetError() error {
	return this.err
}
