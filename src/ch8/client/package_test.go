package client

import (
	//导入包路径：从GOPATH下的src以后的路径开始写起
	"learngo/src/ch8/series" //本项目的包 {module}/{path}
	"testing"
)

func TestPkg(t *testing.T) {
	t.Log(series.GetFib(6))
	t.Log(series.Square(3))
}
