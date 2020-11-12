package maps

import (
	"strconv"
	"sync"
	"testing"

	cmap "github.com/streamrail/concurrent-map"
)

const (
	NumOfReader = 10
	NumOfWrite  = 100
)

type Map interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	Remove(key string)
}

func benchmarkMap(b *testing.B, hm Map) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < NumOfWrite; i++ {
					hm.Set(strconv.Itoa(i), i*i)
					hm.Set(strconv.Itoa(i), i*i)
					hm.Remove(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < NumOfReader; i++ {
					hm.Get(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncMap(b *testing.B) {
	b.Run("map with RWLock", func(b *testing.B) {
		hm := CreateRWLockMap()
		benchmarkMap(b, hm)
	})

	b.Run("map with SyncMap", func(b *testing.B) {
		hm := CreateSyncMapBenchmarkAdapter()
		benchmarkMap(b, hm)
	})

	b.Run("map with ConcurrentMap", func(b *testing.B) {
		hm := cmap.New()
		benchmarkMap(b, &hm)
	})
}
