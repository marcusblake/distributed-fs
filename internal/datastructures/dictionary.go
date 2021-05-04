package datastructures

import (
	"sync"
)

type ConcurrentDictionary struct {
	lock        sync.RWMutex
	internalMap map[interface{}]interface{}
}

func NewDictionary() *ConcurrentDictionary {
	return &ConcurrentDictionary{
		lock:        sync.RWMutex{},
		internalMap: make(map[interface{}]interface{}),
	}
}

func (dict *ConcurrentDictionary) Get(key interface{}) interface{} {
	dict.lock.RLock()
	defer dict.lock.RUnlock()

	if value, ok := dict.internalMap[key]; ok {
		return value
	} else {
		return nil
	}
}

func (dict *ConcurrentDictionary) Contains(key interface{}) bool {
	dict.lock.RLock()
	defer dict.lock.RUnlock()

	_, ok := dict.internalMap[key]
	return ok
}

func (dict *ConcurrentDictionary) Put(key interface{}, value interface{}) {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	dict.internalMap[key] = value
}

func (dict *ConcurrentDictionary) Size() int {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	return len(dict.internalMap)
}
