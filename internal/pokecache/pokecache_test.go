package pokecache

import (
	"testing"
	"fmt"
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

func TestGetNonExistent(t *testing.T) {
	cache := NewCache(5 * time.Second)
	_, ok := cache.Get("https://doesnotexist.com")
	if ok {
		t.Errorf("expected key not exists")
	}
}

func TestOverwrite(t *testing.T) {
    cache := NewCache(5 * time.Second)
    cache.Add("key1", []byte("original"))
    cache.Add("key1", []byte("updated"))
    
    val, ok := cache.Get("key1")
    if !ok {
        t.Errorf("expected to find key")
    }
    if string(val) != "updated" {
        t.Errorf("expected updated value, got %s", string(val))
    }
}
