package main

import (
	"time"
)

func main() {
	mutex := NewMapMutex()
	channel := NewMapChannel()

	for i := 0; i < 10; i++ {
		mutex.Write()
		channel.Write()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for {
				mutex.Read()
			}
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for {
				channel.Read()
			}
		}()
	}

	time.Sleep(1 * time.Second)
}
