package cache

import (
	"errors"
	"sync"
)

type Provider struct {
	HoldMap map[string]string
	Mutex   *sync.Mutex
}

func NewCacheProvider() *Provider {
	return &Provider{
		HoldMap: make(map[string]string),
		Mutex:   &sync.Mutex{},
	}
}

func (p *Provider) SetKey(key string, val string) error {
	p.Mutex.Lock()
	p.HoldMap[key] = val
	p.Mutex.Unlock()
	return nil
}

func (p *Provider) GetKey(key string) (string, error) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	if val, ok := p.HoldMap[key]; ok {
		return val, nil
	}
	return "", errors.New("key not found!")
}
