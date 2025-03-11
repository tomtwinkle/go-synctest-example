package example01

import (
	"errors"
	"sync"
)

var (
	ErrNotFound    = errors.New("not found")
	ErrInvalidData = errors.New("invalid data")
)

type Pool struct {
	pool *sync.Map
}

func NewPool() *Pool {
	return &Pool{
		pool: new(sync.Map),
	}
}

func (p *Pool) SetFoo(id FooID, v Foo) {
	p.pool.Store(id, v)
}

func (p *Pool) GetFoo(id FooID) (*Foo, error) {
	t, ok := p.pool.Load(id)
	if !ok {
		return nil, ErrNotFound
	}
	v, ok := t.(Foo)
	if !ok {
		return nil, ErrInvalidData
	}
	return &v, nil
}
