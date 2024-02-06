package patterns

import (
	"sync"
)

type singleton struct {
	// Define any fields you need for your singleton
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() *singleton {
	// Use sync.Once to ensure that the initialization
	// function is called only once, even in concurrent
	// scenarios.
	once.Do(func() {
		instance = &singleton{} // Create the singleton instance
	})
	return instance
}

// Example usage:
// s := singleton.GetInstance()
