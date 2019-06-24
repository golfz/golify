package golify

type golifyObject struct {
    Value interface{}
    Err   *golifyErr
}

type golifyErr struct {
    ErrCode    int
    ErrMessage string
}

func Init(value interface{}) golifyObject {
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
                ErrCode:errCode,
                ErrMessage:errMsg,
            },
        }
    }

    return golifyObject{
        Value: g.Value,
        Err: nil,
    }
}
