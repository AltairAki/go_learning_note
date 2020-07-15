package remote

import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

/*
	使用远程包
	go get -u github.com/easierway/concurrent_map

	自己包提交到 github 的时候不要带 src 提交
*/
func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
