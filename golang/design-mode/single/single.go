package main

import "sync"

// 单例模式

type single struct {
}

var lock sync.Mutex
var ins *single

func GetIns() *single {
	if ins == nil {
		lock.Lock()
		defer lock.Unlock()
		// 防止在多协程场景下 实例化多个对象
		if ins == nil {
			ins = &single{}
		}
	}

	return ins
}
