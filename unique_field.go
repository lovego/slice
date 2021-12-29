package slice

import (
	"reflect"

	valuePkg "github.com/lovego/value"
)

// UniqueField returns unique values of field from a struct slice
func UniqueField(slice interface{}, fieldPaths ...string) (result []interface{}) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	for i := 0; i < length; i++ {
		value := valuePkg.Get(sliceValue.Index(i), fieldPaths)
		v := value.Interface()
		if !ContainsInterface(result, v) {
			result = append(result, v)
		}
	}
	return
}

// UniqueFieldString returns unique values of field from a struct slice
func UniqueFieldString(slice interface{}, fieldPaths ...string) (result []string) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	for i := 0; i < length; i++ {
		value := valuePkg.Get(sliceValue.Index(i), fieldPaths)
		v := value.String()
		if !ContainsString(result, v) {
			result = append(result, v)
		}
	}
	return
}

// UniqueFieldBool returns unique values of field from a struct slice
func UniqueFieldBool(slice interface{}, fieldPaths ...string) (result []bool) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	for i := 0; i < length; i++ {
		value := valuePkg.Get(sliceValue.Index(i), fieldPaths)
		v := value.Bool()
		if !ContainsBool(result, v) {
			result = append(result, v)
		}
	}
	return
}

// UniqueFieldInt returns unique values of field from a struct slice
func UniqueFieldInt(slice interface{}, fieldPaths ...string) (result []int) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	for i := 0; i < length; i++ {
		value := valuePkg.Get(sliceValue.Index(i), fieldPaths)
		v := int(value.Int())
		if !ContainsInt(result, v) {
			result = append(result, v)
		}
	}
	return
}

// UniqueFieldInt8 returns unique values of field from a struct slice
func UniqueFieldInt8(slice interface{}, fieldPaths ...string) (result []int8) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	for i := 0; i < length; i++ {
		value := valuePkg.Get(sliceValue.Index(i), fieldPaths)
		v := int8(value.Int())
		if !ContainsInt8(result, v) {
			result = append(result, v)
		}
	}
	return
}

// UniqueFieldInt16 returns unique values of field from a struct slice
func UniqueFieldInt16(slice interface{}, fieldPaths ...string) (result []int16) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	for i := 0; i < length; i++ {
		value := valuePkg.Get(sliceValue.Index(i), fieldPaths)
		v := int16(value.Int())
		if !ContainsInt16(result, v) {
			result = append(result, v)
		}
	}
	return
}

// UniqueFieldInt32 returns unique values of field from a struct slice
func UniqueFieldInt32(slice interface{}, fieldPaths ...string) (result []int32) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	for i := 0; i < length; i++ {
		value := valuePkg.Get(sliceValue.Index(i), fieldPaths)
		v := int32(value.Int())
		if !ContainsInt32(result, v) {
			result = append(result, v)
		}
	}
	return
}

// UniqueFieldInt64 returns unique values of field from a struct slice
func UniqueFieldInt64(slice interface{}, fieldPaths ...string) (result []int64) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	for i := 0; i < length; i++ {
		value := valuePkg.Get(sliceValue.Index(i), fieldPaths)
		v := value.Int()
		if !ContainsInt64(result, v) {
			result = append(result, v)
		}
	}
	return
}

// UniqueFieldUint returns unique values of field from a struct slice
func UniqueFieldUint(slice interface{}, fieldPaths ...string) (result []uint) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	for i := 0; i < length; i++ {
		value := valuePkg.Get(sliceValue.Index(i), fieldPaths)
		v := uint(value.Uint())
		if !ContainsUint(result, v) {
			result = append(result, v)
		}
	}
	return
}

// UniqueFieldUint8 returns unique values of field from a struct slice
func UniqueFieldUint8(slice interface{}, fieldPaths ...string) (result []uint8) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	for i := 0; i < length; i++ {
		value := valuePkg.Get(sliceValue.Index(i), fieldPaths)
		v := uint8(value.Uint())
		if !ContainsUint8(result, v) {
			result = append(result, v)
		}
	}
	return
}

// UniqueFieldUint16 returns unique values of field from a struct slice
func UniqueFieldUint16(slice interface{}, fieldPaths ...string) (result []uint16) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	for i := 0; i < length; i++ {
		value := valuePkg.Get(sliceValue.Index(i), fieldPaths)
		v := uint16(value.Uint())
		if !ContainsUint16(result, v) {
			result = append(result, v)
		}
	}
	return
}

// UniqueFieldUint32 returns unique values of field from a struct slice
func UniqueFieldUint32(slice interface{}, fieldPaths ...string) (result []uint32) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	for i := 0; i < length; i++ {
		value := valuePkg.Get(sliceValue.Index(i), fieldPaths)
		v := uint32(value.Uint())
		if !ContainsUint32(result, v) {
			result = append(result, v)
		}
	}
	return
}

// UniqueFieldUint64 returns unique values of field from a struct slice
func UniqueFieldUint64(slice interface{}, fieldPaths ...string) (result []uint64) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	for i := 0; i < length; i++ {
		value := valuePkg.Get(sliceValue.Index(i), fieldPaths)
		v := uint64(value.Uint())
		if !ContainsUint64(result, v) {
			result = append(result, v)
		}
	}
	return
}
