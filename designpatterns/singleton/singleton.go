package singleton

import "sync"

var once sync.Once

type single struct {
}

var Singleton *single

func getSingleton() *single {
	if Singleton == nil {
		once.Do(
			func() {
				Singleton = &single{}
			},
		)
	}
	return Singleton
}
