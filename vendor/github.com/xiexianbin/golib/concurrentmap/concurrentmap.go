// Copyright 2020 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// https://github.com/xiexianbin/gsync/blob/7431fdacf32cf56ea2896abe121a7850f5d8ad3b/utils/concurrentmap.go

package concurrentmap

import (
	"sync"
)

type ConcurrentMap []*Shared

// ShareCount default share count
const ShareCount int = 24

// Shared map share
type Shared struct {
	Items map[string]interface{} // current shared map
	Mu    sync.RWMutex           // current shared mutex
}

// New create new ConcurrentMap
func New() *ConcurrentMap {
	m := make(ConcurrentMap, ShareCount)
	for i := 0; i < ShareCount; i++ {
		m[i] = &Shared{
			Items: map[string]interface{}{},
		}
	}
	return &m
}

// GetSharedMap get special shared map by key
func (m ConcurrentMap) GetSharedMap(key string) *Shared {
	return m[uint(fnv32(key))%uint(ShareCount)]
}

// fnv32 hash func
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	prime32 := uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}

	return hash
}

// Set key and value
func (m ConcurrentMap) Set(key string, value interface{}) {
	sharedMap := m.GetSharedMap(key) // get special shared map by key
	sharedMap.Mu.Lock()              // lock
	sharedMap.Items[key] = value     // set value
	sharedMap.Mu.Unlock()            // unlock
}

// Get value by key
func (m ConcurrentMap) Get(key string) (value interface{}, ok bool) {
	sharedMap := m.GetSharedMap(key) // get special shared map by key
	sharedMap.Mu.RLock()             // read lock
	value, ok = sharedMap.Items[key] // get value
	sharedMap.Mu.RUnlock()           // unlock
	return value, ok
}

// Count get keys sum
func (m ConcurrentMap) Count() int {
	count := 0
	for i := 0; i < ShareCount; i++ {
		m[i].Mu.RLock()          // read lock
		count += len(m[i].Items) // read count
		m[i].Mu.RUnlock()        // unlock
	}
	return count
}

// Keys use goroutine get all keys
func (m ConcurrentMap) Keys() []string {
	count := m.Count()
	keys := make([]string, count)
	chs := make(chan string, count)

	// new goroutine
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(ShareCount)

		for i := 0; i < ShareCount; i++ {
			// pre shared map, new goroutine to statics
			go func(ms *Shared) {
				defer wg.Done()

				ms.Mu.RLock() // read locak
				for k := range ms.Items {
					chs <- k // put key to chs
				}
				ms.Mu.RUnlock() // unlock
			}(m[i])
		}

		// wait all goroutine stop
		wg.Wait()
		close(chs) // close chs, if not next range never stop
	}()

	// circle ch, put all shared key to keys
	for k := range chs {
		keys = append(keys, k)
	}
	return keys
}
