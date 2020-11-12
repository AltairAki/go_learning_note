package maps

import "sync"

type SyncMapBenchmarkAdapter struct {
	m sync.Map
}

func (m *SyncMapBenchmarkAdapter) Get(key string) (interface{}, bool) {
	return m.m.Load(key)
}

func (m *SyncMapBenchmarkAdapter) Set(key string, val interface{}) {
	m.m.Store(key, val)
}

func (m *SyncMapBenchmarkAdapter) Remove(key string) {
	m.m.Delete(key)
}

func CreateSyncMapBenchmarkAdapter() *SyncMapBenchmarkAdapter {
	return &SyncMapBenchmarkAdapter{}
}
