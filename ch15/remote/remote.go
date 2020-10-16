package remote

import (
	cm "github.com/easierway/concurrent_map"
)

func main() {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)

	// a = (m.Get(cm.StrKey("key")))
}
