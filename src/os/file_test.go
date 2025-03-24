package os

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"syscall"
	"testing"
)

func TestFile1(t *testing.T) {
	t.Log(os.TempDir())
	dir, _ := os.MkdirTemp("", "Ts*cpq")
	info, err := os.Stat(dir)
	t.Log(dir)
	t.Log(info, err)
}

func TestFile2(t *testing.T) {
	files, err := os.ReadDir("/home/cpq/eCourse")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func TestFile3(t *testing.T) {
	stat, err := os.Stat("/home/cpq/eCourse")
	if err != nil {
		return
	}
	t.Log(stat.Mode())
}

func TestFile4(t *testing.T) {
	t.Log(os.TempDir())
	t.Log(path.Clean(os.TempDir()))
	t.Log(path.Base(os.TempDir()))

	dir := path.Dir(path.Clean(os.TempDir()))
	t.Log(dir)
	base := path.Base(os.TempDir())

	t.Log(filepath.Join(dir, base+"xxx"))
}

func TestFile5(t *testing.T) {
	wd, _ := syscall.Getwd()
	t.Log(wd)
}
