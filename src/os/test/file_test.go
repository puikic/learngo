package test

import (
	"syscall"
	"testing"
)

func TestFile(t *testing.T) {
	wd, _ := syscall.Getwd()
	t.Log(wd)
}
