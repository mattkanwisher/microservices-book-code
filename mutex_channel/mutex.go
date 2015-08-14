package main

import (
	"math/rand"
	"sync"
)

type MapMutex struct {
	mutex *sync.Mutex
	data  map[int]int
}

func NewMapMutex() *MapMutex {
	return &MapMutex{
		mutex: &sync.Mutex{},
		data:  make(map[int]int),
	}
}

func (m *MapMutex) Read() int {
	key := rand.Intn(5)

	m.mutex.Lock()
	val := m.data[key]
	m.mutex.Unlock()

	return val
}

func (m *MapMutex) Write() {
	key := rand.Intn(5)
	val := rand.Intn(100)

	m.mutex.Lock()
	m.data[key] = val
	m.mutex.Unlock()
}
