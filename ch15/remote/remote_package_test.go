package remote

import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

/*
	使用远程包
	go get -u github.com/easierway/concurrent_map
*/
func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
