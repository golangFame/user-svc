package cache

import (
	"time"
)

type Services interface {
	// Get retrieves an item from the cache. Returns the item or nil, and a bool indicating
	Get(key string, value interface{}) error

	// Set sets an item to the cache, replacing any existing item.
	Set(key string, value interface{}, expire time.Duration) error

	// Delete removes an item from the cache. Does nothing if the key is not in the cache.
	Delete(key string) error

	// Flush seletes all items from the cache.
	Flush() error
}
