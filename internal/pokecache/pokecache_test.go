package pokecache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestCache_ConcurrentAdd(t *testing.T) {
	// A short reap interval to ensure contenction with reaping
	reapInterval := 10 * time.Millisecond
	cache := NewCache(reapInterval)

	numGoroutines := 100
	writesPerGoroutines := 100
	var wg sync.WaitGroup

	t.Logf("Running %d goroutines, each adding %d entries concurrently...", numGoroutines, writesPerGoroutines)

	startSignal := make(chan struct{})

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(gID int) {
			defer wg.Done()
			<-startSignal
			for j := 0; j < writesPerGoroutines; j++ {
				key := fmt.Sprintf("goroutine_%d_key_%d", gID, j)
				val := []byte(fmt.Sprintf("value_%d_%d", gID, j))
				cache.Add(key, val)
			}
		}(i)
	}

	// Release all goroutines simultaneously
	close(startSignal)

	wg.Wait()

	t.Logf("All Add operations completed. Checking cache size...")

	// At this point, the map should contain exactly numGoroutines * writesPerGoroutine entries.
	// If locking is incorrect, concurrent writes might overwrite each other or cause panics.
	// We need to lock the mutex to safely check the map size.
	cache.Lock()
	actualSize := len(cache.entries)
	cache.Unlock()

	expectedSize := numGoroutines * writesPerGoroutines
	if actualSize != expectedSize {
		t.Errorf("Concurrent Add test failed! Expected cache size %d, got %d. This strongly indicates race conditions", expectedSize, actualSize)
	} else {
		t.Logf("Concurrent Add test passed: Cache size %d matches expected size. Locking appears correct.", actualSize)
	}

	// Verify some random entries
	t.Log("Verifying some random entries...")
	for i := 0; i < 5; i++ { // Check 5 random entries
		gID := i % numGoroutines
		j := i % writesPerGoroutines
		key := fmt.Sprintf("goroutine_%d_key_%d", gID, j)
		expectedVal := []byte(fmt.Sprintf("value_%d_%d", gID, j))

		val, ok := cache.Get(key)
		if !ok {
			t.Errorf("Failed to retrieve key %s which should exist.", key)
		} else if string(val) != string(expectedVal) {
			t.Errorf("Value for key %s mismatch. Expected %s, got %s. Race condition possible.", key, string(expectedVal), string(val))
		}
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
