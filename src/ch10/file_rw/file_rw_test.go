package file_rw

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestFileRw(t *testing.T) {
	fin, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()
	fout, err := os.OpenFile("b.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()
	buf := make([]byte, 1024)
	for {
		n, err := fin.Read(buf)
		if err != nil {
			if err == io.EOF {
				if n > 0 {
					fout.Write(buf[:n])
				}
			} else {
				log.Fatal(err)
			}
			break
		} else {
			fout.Write(buf[:n])
		}
	}
}
