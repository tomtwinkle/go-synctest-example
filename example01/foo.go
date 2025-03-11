package example01

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
	f := NewFoo(*id)
	s.pool.SetFoo(*id, *f)
	return f, nil
}
