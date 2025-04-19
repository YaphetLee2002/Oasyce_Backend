package di

import (
	"reflect"

	"go.uber.org/fx"
)

// StructConstructor returns a constructor of the struct type of `structPointer` with all the exported fields of it.
// Notice: the injected fields should be in UPPERCASE, since we used reflect to get the fields.
//
// Example:
//
// Suppose we have a struct `S` which contains two dependencies:
//
//	type S struct {
//	  F1 S1
//	  F2 S2
//	}
//
// then
//
//	StructConstructor(new(S))
//
// returns:
//
//	func(f1 S1, f2 S2) *S {
//	  return &S{
//		 F1: f1,
//	    F2: f2,
//	  }
//	}
//
// When used in the case of dependency injection, it is often convenient to omit trivial constructor, i.e. instead of writing something like this
//
//	func NewS(f1 S1, f2 S2) *S {
//	  return &S{
//		 F1: f1,
//	    F2: f2,
//	  }
//	}
//	fx.Provide(NewS)
//
// you can use:
//
//	fx.Provide(StructConstructor(new(S)))
func StructConstructor(structPointer interface{}) interface{} {
	structPointerType := reflect.TypeOf(structPointer)
	if structPointerType.Kind() != reflect.Pointer || structPointerType.Elem().Kind() != reflect.Struct {
		panic("only support struct pointer")
	}
	structType := structPointerType.Elem()
	numField := structType.NumField()
	var in []reflect.Type
	var tags []string
	hasTag := false
	for i := 0; i < numField; i++ {
		field := structType.Field(i)
		if !field.IsExported() {
			panic("provided struct contains unexported field(s)")
		}
		in = append(in, field.Type)
		tags = append(tags, string(field.Tag))
		if len(string(field.Tag)) > 0 {
			hasTag = true
		}
	}
	out := []reflect.Type{structPointerType}
	funcType := reflect.FuncOf(in, out, false)

	rtn := reflect.MakeFunc(funcType, func(args []reflect.Value) (results []reflect.Value) {
		resultValue := reflect.New(structType)
		for i := 0; i < numField; i++ {
			fieldValue := resultValue.Elem().Field(i)
			fieldValue.Set(args[i])
		}
		return []reflect.Value{resultValue}
	}).Interface()

	if hasTag {
		rtn = fx.Annotate(rtn, fx.ParamTags(tags...))
	}

	return rtn
}
