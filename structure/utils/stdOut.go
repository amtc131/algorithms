package utils

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func NewStDoutRestore(t *testing.T) (testStdout, writer, osStdout *os.File) {
	testStdout, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() err = %v; want %v", err, nil)
	}
	osStdout = os.Stdout
	os.Stdout = writer

	return testStdout, writer, osStdout
}

func Stdout(writer, testStdout *os.File) (got string) {
	writer.Close()
	var buf bytes.Buffer
	io.Copy(&buf, testStdout)
	got = buf.String()
	return got
}
