package singleton_lazy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLazySingleton(t *testing.T) {
	assert.Equal(t, GetLazySingleton(), GetLazySingleton())
}

func BenchmarkGetLazySingleton(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazySingleton() != GetLazySingleton() {
				b.Errorf("test fail")
			}
		}
	})
}
