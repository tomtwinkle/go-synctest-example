package example01

import (
	"time"

	"github.com/google/uuid"
)

type FooID uuid.UUID

func NewFooID() (*FooID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return (*FooID)(&id), nil
}

func (id FooID) UUID() uuid.UUID {
	return uuid.UUID(id)
}

type Foo struct {
	ID        FooID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewFoo(id FooID) *Foo {
	return &Foo{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
