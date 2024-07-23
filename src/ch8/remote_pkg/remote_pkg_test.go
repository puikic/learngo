package remote_pkg

import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

func TestRemotePkg(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}