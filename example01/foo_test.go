package example01_test

import (
	"math/rand/v2"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/uuid"
	"github.com/tomtwinkle/go-synctest-example/example01"
)

// Sleep for a random number of minutes and change to the new time.
// Must used in synctest.Run.
func changeRandTime(now time.Time) time.Time {
	d := rand.N(time.Minute)
	time.Sleep(d)
	return now.Add(d)
}

// synctest is an experimental GOEXPERIMENT=synctest must be specified as an environment variable.
func TestFooService_CreateFoo(t *testing.T) {
	s := example01.NewFooService()

	t.Parallel()
	t.Run("create foo using synctest", func(t *testing.T) {
		t.Parallel()
		synctest.Run(func() {
			now := time.Now()
			expectedTime := changeRandTime(now)

			got, err := s.CreateFoo()
			if err != nil {
				t.Errorf("want error: nil, got: %+v", err)
			}
			if got == nil {
				t.Errorf("want foo: not nil, got: nil")
			}
			if got.ID.UUID() == uuid.Nil {
				t.Errorf("want foo id: not nil, got: %v", got.ID.UUID())
			}
			if !got.CreatedAt.Equal(expectedTime) {
				t.Errorf("want created at: %v, got: %v", now, got.CreatedAt)
			}
			if !got.UpdatedAt.Equal(expectedTime) {
				t.Errorf("want updated at: %v, got: %v", now, got.UpdatedAt)
			}
		})
	})
}
