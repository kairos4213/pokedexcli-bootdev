package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = time.Second * 5
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "www.api.com/endpnt/1",
			val: []byte("val1"),
		},
		{
			key: "www.api.com/endpnt/2",
			val: []byte("val2"),
		},
		{
			key: "www.api.com/endpnt/3",
			val: []byte("val3"),
		},
	}

	for i, cas := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(cas.key, cas.val)
			val, ok := cache.Get(cas.key)
			if !ok {
				t.Errorf("key: '%v' not found", cas.key)
				return
			}
			if string(val) != string(cas.val) {
				t.Errorf("val: '%v' does not match '%v'", cas.val, val)
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const intervalTime = time.Millisecond * 5
	const waitTime = intervalTime + 5*time.Millisecond
	cache := NewCache(intervalTime)
	cache.Add("www.api.com/endpnt/", []byte("byte"))

	_, ok := cache.Get("www.api.com/endpnt/")
	if !ok {
		t.Errorf("key not found")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("www.api.com/endpnt/")
	if ok {
		t.Errorf("key found -- reaping did not occur")
		return
	}
}
