package fish

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/awnumar/memguard"
)

type Brain struct {
	memories map[string]*secureMemory
	mux      sync.Mutex
}

func NewBrain() (*Brain) {
	brain := &Brain{
		memories: map[string]*secureMemory{},
	}

	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				brain.cleanUp()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return brain
}

func (b *Brain) cleanUp() {
	for k := range b.memories {
		b.Get(k) // Implicitly destroys expired memories
	}
}

func (b *Brain) Get(key string) ([]byte, error) {
	b.mux.Lock()
	defer b.mux.Unlock()

	secMem, ok := b.memories[key]
	if !ok {
		return nil, errors.New("key not present")
	}

	if secMem.expired() {
		fmt.Printf("Destroying expired %s\n", key)
		secMem.destroy()
		delete(b.memories, key)
		return nil, errors.New("key expired")
	}

	if secMem.value != nil {
		return secMem.value.Buffer(), nil
	} else {
		return nil, nil
	}
}

func (b *Brain) Set(key string, val []byte, expiration time.Time) (error) {
	b.mux.Lock()
	defer b.mux.Unlock()

	secureMemory := &secureMemory{
		expiration: expiration,
	}

	if len(val) != 0 {
		lockedBuffer, err := memguard.NewImmutableFromBytes([]byte(val))
		if err != nil {
			return err
		}

		secureMemory.value = lockedBuffer
	}

	b.memories[key] = secureMemory

	return nil
}
