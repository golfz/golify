package golify

type golifyIntegerObject struct {
    Value int64
    Err   *golifyErr
}

//func (g golifyIntegerObject) MoreThan(min interface{}, errCode int, errMsg string) golifyIntegerObject {
//
//    if g.Err != nil {
//        return g
//    }
//
//    fMin := 1.0
//
//    if reflect.TypeOf(min) == reflect.TypeOf(int(1)) {
//        fMin = float64(min.(int))
//
//    } else if reflect.TypeOf(min) == reflect.TypeOf(int64(1)) {
//        fMin = float64(min.(int64))
//
//    } else if reflect.TypeOf(min) == reflect.TypeOf(1.1) {
//        fMin = min.(float64)
//
//    } else {
//        return golifyUnknownObject{
//            Value: g.Value,
//            Err: &golifyErr{
//                ErrCode: errCode,
//                ErrMsg:  errMsg,
//            },
//        }
//    }
//
//    gVal := 1.0
//
//    if reflect.TypeOf(g.Value) == reflect.TypeOf(1) {
//        gVal = float64(g.Value.(int))
//
//    } else if reflect.TypeOf(g.Value) == reflect.TypeOf(int64(1)) {
//        gVal = float64(g.Value.(int64))
//
//    } else if reflect.TypeOf(g.Value) == reflect.TypeOf(1.1) {
//        gVal = g.Value.(float64)
//
//    } else {
//        return golifyUnknownObject{
//            Value: g.Value,
//            Err: &golifyErr{
//                ErrCode: errCode,
//                ErrMsg:  errMsg,
//            },
//        }
//    }
//
//    if gVal < fMin {
//        return golifyUnknownObject{
//            Value: g.Value,
//            Err: &golifyErr{
//                ErrCode: errCode,
//                ErrMsg:  errMsg,
//            },
//        }
//    }
//
//    return golifyUnknownObject{
//        Value: g.Value,
//        Err:   nil,
//    }
//}
//
//func (g golifyIntegerObject) LessThan(max interface{}, errCode int, errMsg string) golifyUnknownObject {
//
//    if g.Err != nil {
//        return g
//    }
//
//    fMax := 1.0
//
//    if reflect.TypeOf(max) == reflect.TypeOf(1) {
//        fMax = float64(max.(int))
//
//    } else if reflect.TypeOf(max) == reflect.TypeOf(int64(1)) {
//        fMax = float64(max.(int64))
//
//    } else if reflect.TypeOf(max) == reflect.TypeOf(1.1) {
//        fMax = max.(float64)
//
//    } else {
//        return golifyUnknownObject{
//            Value: g.Value,
//            Err: &golifyErr{
//                ErrCode: errCode,
//                ErrMsg:  errMsg,
//            },
//        }
//    }
//
//    gVal := 1.0
//
//    if reflect.TypeOf(g.Value) == reflect.TypeOf(1) {
//        gVal = float64(g.Value.(int))
//
//    } else if reflect.TypeOf(g.Value) == reflect.TypeOf(int64(1)) {
//        gVal = float64(g.Value.(int64))
//
//    } else if reflect.TypeOf(g.Value) == reflect.TypeOf(1.1) {
//        gVal = g.Value.(float64)
//
//    } else {
//        return golifyUnknownObject{
//            Value: g.Value,
//            Err: &golifyErr{
//                ErrCode: errCode,
//                ErrMsg:  errMsg,
//            },
//        }
//    }
//
//    if gVal > fMax {
//        return golifyUnknownObject{
//            Value: g.Value,
//            Err: &golifyErr{
//                ErrCode: errCode,
//                ErrMsg:  errMsg,
//            },
//        }
//    }
//
//    return golifyUnknownObject{
//        Value: g.Value,
//        Err:   nil,
//    }
//}
