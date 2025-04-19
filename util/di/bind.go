package di

import (
	"reflect"
)

// Bind returns a constructor which use type of `structValue` to implement the interface type of `interfaceValue`.
//
// Example:
//
// Suppose struct `S` implements interface `I`, then:
//
//	Bind(new(S), new(I))
//
// returns:
//
//	func(s *S) I { return s }
//
// When used with dependency injection, Bind means provide type `I` which implemented by type `S`, i.e.
//
//	fx.Provide(Bind(new(s), new(I)))
//
// equals to:
//
//	fx.Provide(func(s *S) I { return s })
func Bind(structValue, interfaceValue interface{}) interface{} {
	interfaceType := reflect.TypeOf(interfaceValue)
	if interfaceType.Kind() == reflect.Ptr {
		interfaceType = interfaceType.Elem()
	}
	if interfaceType.Kind() != reflect.Interface {
		panic("the type of parameter 'interfaceValue' should be interface")
	}

	in := []reflect.Type{reflect.TypeOf(structValue)}
	out := []reflect.Type{interfaceType}
	funcType := reflect.FuncOf(in, out, false)

	return reflect.MakeFunc(funcType, func(args []reflect.Value) (results []reflect.Value) {
		return args
	}).Interface()
}
