package golify

import (
    "reflect"
    "strconv"
)

type golifyObject struct {
    Value interface{}
    Err   *golifyErr
}

type golifyErr struct {
    ErrCode int
    ErrMsg  string
}

func Validate(value interface{}) golifyObject {
    return golifyObject{
        Value: value,
        Err:   nil,
    }
}

func (g golifyObject) NotNil(errCode int, errMsg string) golifyObject {

    if g.Err != nil {
        return g
    }

    if g.Value == nil {
        return golifyObject{
            Value: g.Value,
            Err: &golifyErr{
                ErrCode: errCode,
                ErrMsg:  errMsg,
            },
        }
    }

    return golifyObject{
        Value: g.Value,
        Err:   nil,
    }
}

func (g golifyObject) IsString(errCode int, errMsg string) golifyObject {

    if g.Err != nil {
        return g
    }

    if reflect.TypeOf(g.Value) != reflect.TypeOf("abc") {
        return golifyObject{
            Value: g.Value,
            Err: &golifyErr{
                ErrCode: errCode,
                ErrMsg:  errMsg,
            },
        }
    }

    return golifyObject{
        Value: g.Value,
        Err:   nil,
    }
}

func (g golifyObject) IsInt(errCode int, errMsg string) golifyObject {

    if g.Err != nil {
        return g
    }

    if reflect.TypeOf(g.Value) != reflect.TypeOf(1) {
        return golifyObject{
            Value: g.Value,
            Err: &golifyErr{
                ErrCode: errCode,
                ErrMsg:  errMsg,
            },
        }
    }

    return golifyObject{
        Value: g.Value,
        Err:   nil,
    }
}

func (g golifyObject) MoreThan(min interface{}, errCode int, errMsg string) golifyObject {

    if g.Err != nil {
        return g
    }

    fMin := 1.0

    if reflect.TypeOf(min) == reflect.TypeOf(1) {
        fMin = float64(min.(int))

    } else if reflect.TypeOf(min) == reflect.TypeOf(1.1) {
        fMin = min.(float64)

    } else {
        return golifyObject{
            Value: g.Value,
            Err: &golifyErr{
                ErrCode: errCode,
                ErrMsg:  errMsg,
            },
        }
    }

    gVal := 1.0

    if reflect.TypeOf(g.Value) == reflect.TypeOf(1) {
        gVal = float64(g.Value.(int))

    } else if reflect.TypeOf(g.Value) == reflect.TypeOf(1.1) {
        gVal = g.Value.(float64)

    } else {
        return golifyObject{
            Value: g.Value,
            Err: &golifyErr{
                ErrCode: errCode,
                ErrMsg:  errMsg,
            },
        }
    }

    if gVal < fMin {
        return golifyObject{
            Value: g.Value,
            Err: &golifyErr{
                ErrCode: errCode,
                ErrMsg:  errMsg,
            },
        }
    }

    return golifyObject{
        Value: g.Value,
        Err:   nil,
    }
}

func (g golifyObject) LessThan(max interface{}, errCode int, errMsg string) golifyObject {

    if g.Err != nil {
        return g
    }

    fMax := 1.0

    if reflect.TypeOf(max) == reflect.TypeOf(1) {
        fMax = float64(max.(int))

    } else if reflect.TypeOf(max) == reflect.TypeOf(1.1) {
        fMax = max.(float64)

    } else {
        return golifyObject{
            Value: g.Value,
            Err: &golifyErr{
                ErrCode: errCode,
                ErrMsg:  errMsg,
            },
        }
    }

    gVal := 1.0

    if reflect.TypeOf(g.Value) == reflect.TypeOf(1) {
        gVal = float64(g.Value.(int))

    } else if reflect.TypeOf(g.Value) == reflect.TypeOf(1.1) {
        gVal = g.Value.(float64)

    } else {
        return golifyObject{
            Value: g.Value,
            Err: &golifyErr{
                ErrCode: errCode,
                ErrMsg:  errMsg,
            },
        }
    }

    if gVal > fMax {
        return golifyObject{
            Value: g.Value,
            Err: &golifyErr{
                ErrCode: errCode,
                ErrMsg:  errMsg,
            },
        }
    }

    return golifyObject{
        Value: g.Value,
        Err:   nil,
    }
}

func (g golifyObject) Atoi(errCode int, errMsg string) golifyObject {
    if g.Err != nil {
        return g
    }

    if reflect.TypeOf(g.Value) == reflect.TypeOf("abc") {
        i, e := strconv.Atoi(g.Value.(string))
        if e != nil {
            return golifyObject{
                Value: g.Value,
                Err: &golifyErr{
                    ErrCode: errCode,
                    ErrMsg:  errMsg,
                },
            }
        }

        return golifyObject{
            Value: i,
            Err:   nil,
        }
    }

    return golifyObject{
        Value: g.Value,
        Err: &golifyErr{
            ErrCode: errCode,
            ErrMsg:  errMsg,
        },
    }
}
