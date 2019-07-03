package golify

type golifyStringObject struct {
    Value string
    Err   *golifyErr
}

/*
func (g golifyUnknownObject) Atoi(errCode int, errMsg string) golifyUnknownObject {
    if g.Err != nil {
        return g
    }

    if reflect.TypeOf(g.Value) == reflect.TypeOf("abc") {
        i, e := strconv.Atoi(g.Value.(string))
        if e != nil {
            return golifyUnknownObject{
                Value: g.Value,
                Err: &golifyErr{
                    ErrCode: errCode,
                    ErrMsg:  errMsg,
                },
            }
        }

        return golifyUnknownObject{
            Value: i,
            Err:   nil,
        }
    }

    return golifyUnknownObject{
        Value: g.Value,
        Err: &golifyErr{
            ErrCode: errCode,
            ErrMsg:  errMsg,
        },
    }
}
*/
