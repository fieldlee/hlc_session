package model

import (
	"time"
)

type Token struct {
	U		string
	Left		int64
}

var Logined = make(map[string]*Token)

func (this *Token)Count(key string, left int64) {
	for {
		if this.Left >= left {
			delete(Logined, key)
			return
		}
		time.Sleep(time.Second)
		this.Left++
	}
}
