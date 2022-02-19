package test

import (
	"github.com/beyzaekici/getir-case-study/data/store/cache"
	"testing"
)

func TestGetMemory(t *testing.T) {
	provider := cache.NewCacheProvider()

	err := provider.SetKey("key", "value")
	if err != nil {
		return
	}
	v, e := provider.GetKey("key")
	if e != nil {
		t.Fail()
	}

	if v != "value" {
		t.Fail()
	}
}

func TestSetMemory(t *testing.T) {
	provider := cache.NewCacheProvider()
	err := provider.SetKey("key", "value")
	if err != nil {
		return
	}
	if k, ok := provider.HoldMap["key"]; !ok {
		t.Fail()
	} else if k != "value" {
		t.Fail()
	}
}
