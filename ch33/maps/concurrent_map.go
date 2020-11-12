package maps

import "github.com/easierway/concurrent_map"

type ConcurrentMapBenckmarkAdapter struct {
	m *concurrent_map.ConcurrentMap
}

func (m *ConcurrentMapBenckmarkAdapter) Get(key interface{}) (interface{}, bool) {
	return m.m.Get(concurrent_map.StrKey(key.(string)))
}

func (m *ConcurrentMapBenckmarkAdapter) Set(key interface{}, val interface{}) {
	m.m.Set(concurrent_map.StrKey(key.(string)), val)
}

func (m *ConcurrentMapBenckmarkAdapter) Remove(key interface{}) {
	m.m.Del(concurrent_map.StrKey(key.(string)))
}

func CreateConcurrentMapBenckmarkAdapter(numOfPartitions int) *ConcurrentMapBenckmarkAdapter {
	concurrentMap := concurrent_map.CreateConcurrentMap(numOfPartitions)
	return &ConcurrentMapBenckmarkAdapter{concurrentMap}
}
