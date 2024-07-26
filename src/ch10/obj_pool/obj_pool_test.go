package obj_pool

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ResuableObj struct{}
type ObjPool struct {
	buf chan *ResuableObj
}

func NewObjPool(num int) *ObjPool {
	objPool := &ObjPool{make(chan *ResuableObj, num)}
	for i := 0; i < num; i++ {
		objPool.buf <- &ResuableObj{}
	}
	return objPool
}
func (p *ObjPool) Get(timeout time.Duration) (*ResuableObj, error) {
	select {
	case obj := <-p.buf:
		return obj, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}
func (p *ObjPool) Put(obj *ResuableObj) error {
	select {
	case p.buf <- obj:
		return nil
	default:
		return errors.New("timeout")
	}
}
func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.Get(time.Second); err != nil {
			t.Error(err)
		} else {
			fmt.Println(v, i)

		}
	}
	fmt.Println("Done!")
}
