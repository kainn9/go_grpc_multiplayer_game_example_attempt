package util

import (
	"math/rand"
	"reflect"
	"time"
)

// returns random number from 0 -> n - 1
func RandomInt(n int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(n)
}

func ToInterfacePtr(val interface{}) interface{} {
	ptr := reflect.New(reflect.TypeOf(val))
	ptr.Elem().Set(reflect.ValueOf(val))
	return ptr.Interface()
}
