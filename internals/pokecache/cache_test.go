package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestEmptyCache(t *testing.T) {
	const interval = 5 * time.Second

	cache := NewCache(interval)
	count := 0

	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "one",
			value: []byte("test"),
		},
		{
			key:   "two",
			value: []byte("another case"),
		},
	}

	for range cache.data {
		count++
	}

	if count > 0 {
		t.Errorf("Expect empty cache when initialized")
		return
	}

	for _, data := range cases {
		cache.Add(data.key, data.value)
		count++
	}

	if count != 2 {
		t.Errorf("Expect empty cache to have 2 elements")
		return
	}
}

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
