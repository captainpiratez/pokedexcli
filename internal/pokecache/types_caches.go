package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}
