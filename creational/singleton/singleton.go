package singleton

import "sync"

type Singleton interface {
	getAddr()
}

type singleton struct {
	Singleton
}

func (s singleton) getAddr() {
	println(&s)
}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
