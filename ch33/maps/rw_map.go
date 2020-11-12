package maps

import "sync"

type RWLockMap struct {
	m    map[string]interface{}
	lock sync.RWMutex
}

func (m *RWLockMap) Get(key string) (interface{}, bool) {
	m.lock.RLock()
	v, ok := m.m[key]
	m.lock.RUnlock()
	return v, ok
}

func (m *RWLockMap) Set(key string, val interface{}) {
	m.lock.Lock()
	m.m[key] = val
	m.lock.Unlock()
}

func (m *RWLockMap) Remove(key string) {
	m.lock.Lock()
	delete(m.m, key)
	m.lock.Unlock()
}

func CreateRWLockMap() *RWLockMap {
	m := make(map[string]interface{}, 0)
	return &RWLockMap{m: m}
}
