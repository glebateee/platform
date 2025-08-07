package services

import (
	"context"
	"reflect"
	"sync"
)

func addTransient(factoryFunc interface{}) error {
	return addService(Transient, factoryFunc)
}

func addScoped(factoryFunc interface{}) error {
	return addService(Scoped, factoryFunc)
}

func addSingleton(factoryFunc interface{}) (err error) {
	factoryFuncVal := reflect.ValueOf(factoryFunc)
	if factoryFuncVal.Kind() == reflect.Func && factoryFuncVal.Type().NumOut() == 1 {
		var results []reflect.Value
		once := sync.Once{}
		wrapper := reflect.MakeFunc(factoryFuncVal.Type(), func([]reflect.Value) []reflect.Value {
			once.Do(func() {
				results = invokeFunction(context.TODO(), factoryFuncVal)
			})
			return results
		})
		err = addService(Singleton, wrapper.Interface())
	}
	return
}
