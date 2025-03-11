package example01

import (
	"time"
)

type FooService struct {
	pool *Pool
}

func NewFooService() *FooService {
	return &FooService{
		pool: NewPool(),
	}
}

func (s *FooService) CreateFoo() (*Foo, error) {
	id, err := NewFooID()
	if err != nil {
		return nil, err
	}
	f := &Foo{
		ID:        *id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s.pool.SetFoo(*id, *f)
	return f, nil
}
