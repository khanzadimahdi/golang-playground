package singletone

import "sync"

var lock = &sync.Mutex{}

type single struct {
	//
}

var singleInstance *single

// way 1:
// We can create a single instance inside the init function.
// This is only applicable if the early initialization of the object is ok.
// The init function is only called once per file in a package,
// so we can be sure that only a single instance will be created.

// func init() {
// 	singleInstance = &single{}
// }

// way 2:
// func GetInstance() *single {
// 	if singleInstance == nil {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		if singleInstance == nil {
// 			singleInstance = &single{}
// 		}
// 	}

// 	return singleInstance
// }

// way 3:
var once = &sync.Once{}

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(func() {
			singleInstance = &single{}
		})
	}

	return singleInstance
}
