package main

import (
	"math/rand"
)

type MapChannel struct {
	lock   chan struct{}
	unlock chan struct{}
	data   map[int]int
}

func NewMapChannel() *MapChannel {
	m := &MapChannel{
		lock:   make(chan struct{}),
		unlock: make(chan struct{}),
		data:   make(map[int]int),
	}

	go func() {
		for {
			select {
			case <-m.lock:
				m.unlock <- struct{}{}
			}
		}
	}()

	return m
}

func (m *MapChannel) Read() int {
	key := rand.Intn(5)
	m.lock <- struct{}{}
	val := m.data[key]
	<-m.unlock

	return val
}

func (m *MapChannel) Write() {
	key := rand.Intn(5)
	val := rand.Intn(100)

	m.lock <- struct{}{}
	m.data[key] = val
	<-m.unlock
}
