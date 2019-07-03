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

    if g.Value == nil {
        return golifyUnknownObject{
            Value: g.Value,
            Err: &golifyErr{
                ErrCode: errCode,
                ErrMsg:  errMsg,
            },
        }
    }

    return golifyUnknownObject{
        Value: g.Value,
        Err:   nil,
    }
}

func (g golifyUnknownObject) ToString(errCode int, errMsg string) golifyStringObject {

    if g.Err != nil {
        return golifyStringObject{
            Value: "",
            Err:   g.Err,
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

func (g golifyUnknownObject) ToInteger(errCode int, errMsg string) golifyIntegerObject {

    if g.Err != nil {
        return golifyIntegerObject{
            Value: 0,
            Err:   g.Err,
        }
    }

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


