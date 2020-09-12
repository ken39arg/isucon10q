package main

import (
	"sync"
)

var (
	cachedata  map[interface{}]interface{}
	cachelock  *sync.RWMutex
	cacheklock map[interface{}]*sync.Mutex
)

func initcach(size int) {
	cachelock = new(sync.RWMutex)
	cachedata = make(map[interface{}]interface{}, size)
	cacheklock = make(map[interface{}]*sync.Mutex, size)
}

func cacheGetOrSet(key interface{}, cb func() (interface{}, error)) (interface{}, error) {
	cachelock.RLock()
	l, ok := cacheklock[key]
	cachelock.RUnlock()
	if !ok {
		cachelock.Lock()
		l = new(sync.Mutex)
		cacheklock[key] = l
		cachelock.Unlock()
	}

	l.Lock()
	defer l.Unlock()
	if v, ok := cachedata[key]; ok {
		return v, nil
	}
	v, err := cb()
	if err != nil {
		return nil, err
	}
	cachedata[key] = v
	return v, nil
}
