package singleton_lazy

import (
	"Design-Pattern/singleton"
	"sync"
)

// 懒汉式 单例模式
var lazySingleton *singleton.Singleton
var once = &sync.Once{}

func GetLazySingleton() *singleton.Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &singleton.Singleton{}
		})
	}
	return lazySingleton
}
