package fish

import (
	"time"

	"github.com/awnumar/memguard"
)

type secureMemory struct {
	value *memguard.LockedBuffer
	expiration time.Time
}

func (sm *secureMemory) expired() bool {
	return sm.expiration.Before(time.Now())
}

func (sm *secureMemory) destroy() {
	sm.value.Destroy()
}