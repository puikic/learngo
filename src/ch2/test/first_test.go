// .go文件必须以_test结尾命名
package try_test

import "testing"

// 方法名字以大写Test开头
func TestFirstTry(t *testing.T) {
	//test
	t.Log("My first try!")
}
