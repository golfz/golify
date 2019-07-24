package golify

import (
	"reflect"
)

type golifyUnknownObject struct {
	Value interface{}
	Err   *golifyErr
}

type golifyErr struct {
	ErrCode int
	ErrMsg  string
}

func Init(value interface{}) golifyUnknownObject {
	return golifyUnknownObject{
		Value: value,
		Err:   nil,
	}
}

func (g golifyUnknownObject) NotNil(errCode int, errMsg string) golifyUnknownObject {

	if g.Err != nil {
		return g
	}

	// BOOLEAN
	b := true
	testBoolPtr := &b
	if reflect.TypeOf(g.Value) == reflect.TypeOf(testBoolPtr) {
		if g.Value.(*bool) == nil {
			return golifyUnknownObject{
				Value: g.Value,
				Err: &golifyErr{
					ErrCode: errCode,
					ErrMsg:  errMsg,
				},
			}
		}
	}

	// STRING
	a := "abc"
	testStrPtr := &a
	if reflect.TypeOf(g.Value) == reflect.TypeOf(testStrPtr) {
		if g.Value.(*string) == nil {
			return golifyUnknownObject{
				Value: g.Value,
				Err: &golifyErr{
					ErrCode: errCode,
					ErrMsg:  errMsg,
				},
			}
		}
	}

	// INTEGER
	i := int(1)
	i8 := int8(1)
	i16 := int16(1)
	i32 := int32(1)
	i64 := int64(1)
	testIntPtr := &i
	testInt8Ptr := &i8
	testInt16Ptr := &i16
	testInt32Ptr := &i32
	testInt64Ptr := &i64

	// pointer
	if reflect.TypeOf(g.Value) == reflect.TypeOf(testIntPtr) {
		if g.Value.(*int) == nil {
			return golifyUnknownObject{
				Value: g.Value,
				Err: &golifyErr{
					ErrCode: errCode,
					ErrMsg:  errMsg,
				},
			}
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf(testInt8Ptr) {
		if g.Value.(*int8) == nil {
			return golifyUnknownObject{
				Value: g.Value,
				Err: &golifyErr{
					ErrCode: errCode,
					ErrMsg:  errMsg,
				},
			}
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf(testInt16Ptr) {
		if g.Value.(*int16) == nil {
			return golifyUnknownObject{
				Value: g.Value,
				Err: &golifyErr{
					ErrCode: errCode,
					ErrMsg:  errMsg,
				},
			}
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf(testInt32Ptr) {
		if g.Value.(*int32) == nil {
			return golifyUnknownObject{
				Value: g.Value,
				Err: &golifyErr{
					ErrCode: errCode,
					ErrMsg:  errMsg,
				},
			}
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf(testInt64Ptr) {
		if g.Value.(*int64) == nil {
			return golifyUnknownObject{
				Value: g.Value,
				Err: &golifyErr{
					ErrCode: errCode,
					ErrMsg:  errMsg,
				},
			}
		}
	}

	return golifyUnknownObject{
		Value: g.Value,
		Err:   nil,
	}
}

func (g golifyUnknownObject) AsString(errCode int, errMsg string) golifyStringObject {

	if g.Err != nil {
		return golifyStringObject{
			Value: "",
			Err:   g.Err,
		}
	}

	a := "abc"
	testStrPtr := &a

	if reflect.TypeOf(g.Value) == reflect.TypeOf(testStrPtr) {
		return golifyStringObject{
			Value: *(g.Value.(*string)),
			Err:   nil,
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf("abc") {
		return golifyStringObject{
			Value: g.Value.(string),
			Err:   nil,
		}
	}

	return golifyStringObject{
		Value: "",
		Err: &golifyErr{
			ErrCode: errCode,
			ErrMsg:  errMsg,
		},
	}

}

func (g golifyUnknownObject) AsInteger(errCode int, errMsg string) golifyIntegerObject {

	if g.Err != nil {
		return golifyIntegerObject{
			Value: 0,
			Err:   g.Err,
		}
	}

	i := int(1)
	i8 := int8(1)
	i16 := int16(1)
	i32 := int32(1)
	i64 := int64(1)
	testIntPtr := &i
	testInt8Ptr := &i8
	testInt16Ptr := &i16
	testInt32Ptr := &i32
	testInt64Ptr := &i64

	// pointer
	if reflect.TypeOf(g.Value) == reflect.TypeOf(testIntPtr) {
		return golifyIntegerObject{
			Value: int64(*g.Value.(*int)),
			Err:   nil,
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf(testInt8Ptr) {
		return golifyIntegerObject{
			Value: int64(*g.Value.(*int8)),
			Err:   nil,
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf(testInt16Ptr) {
		return golifyIntegerObject{
			Value: int64(*g.Value.(*int16)),
			Err:   nil,
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf(testInt32Ptr) {
		return golifyIntegerObject{
			Value: int64(*g.Value.(*int32)),
			Err:   nil,
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf(testInt64Ptr) {
		return golifyIntegerObject{
			Value: int64(*g.Value.(*int64)),
			Err:   nil,
		}
	}

	// non-pointer
	if reflect.TypeOf(g.Value) == reflect.TypeOf(int(1)) {
		return golifyIntegerObject{
			Value: int64(g.Value.(int)),
			Err:   nil,
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf(int8(1)) {
		return golifyIntegerObject{
			Value: int64(g.Value.(int8)),
			Err:   nil,
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf(int16(1)) {
		return golifyIntegerObject{
			Value: int64(g.Value.(int16)),
			Err:   nil,
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf(int32(1)) {
		return golifyIntegerObject{
			Value: int64(g.Value.(int32)),
			Err:   nil,
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf(int64(1)) {
		return golifyIntegerObject{
			Value: g.Value.(int64),
			Err:   nil,
		}
	}

	return golifyIntegerObject{
		Value: 0,
		Err: &golifyErr{
			ErrCode: errCode,
			ErrMsg:  errMsg,
		},
	}

}

func (g golifyUnknownObject) AsBoolean(errCode int, errMsg string) golifyBooleanObject {
	if g.Err != nil {
		return golifyBooleanObject{
			Value: false,
			Err:   g.Err,
		}
	}

	b := true
	testBoolPtr := &b

	if reflect.TypeOf(g.Value) == reflect.TypeOf(testBoolPtr) {
		return golifyBooleanObject{
			Value: *(g.Value.(*bool)),
			Err:   nil,
		}
	}

	if reflect.TypeOf(g.Value) == reflect.TypeOf(true) {
		return golifyBooleanObject{
			Value: g.Value.(bool),
			Err:   nil,
		}
	}

	return golifyBooleanObject{
		Value: false,
		Err: &golifyErr{
			ErrCode: errCode,
			ErrMsg:  errMsg,
		},
	}
}
