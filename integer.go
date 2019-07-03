package golify

type golifyIntegerObject struct {
    Value int64
    Err   *golifyErr
}

func (g golifyIntegerObject) MoreThan(min int64, errCode int, errMsg string) golifyIntegerObject {

    if g.Err != nil {
        return g
    }

    if g.Value > min {
        return golifyIntegerObject{
            Value: g.Value,
            Err:   nil,
        }
    }

    return golifyIntegerObject{
        Value: g.Value,
        Err: &golifyErr{
            ErrCode: errCode,
            ErrMsg:  errMsg,
        },
    }
}

func (g golifyIntegerObject) LessThan(max int64, errCode int, errMsg string) golifyIntegerObject {

    if g.Err != nil {
        return g
    }

    if g.Value < max {
        return golifyIntegerObject{
            Value: g.Value,
            Err:   nil,
        }
    }

    return golifyIntegerObject{
        Value: g.Value,
        Err: &golifyErr{
            ErrCode: errCode,
            ErrMsg:  errMsg,
        },
    }
}
