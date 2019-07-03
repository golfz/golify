package golify

import "strconv"

type golifyStringObject struct {
    Value string
    Err   *golifyErr
}

func (g golifyStringObject) Atoi(errCode int, errMsg string) golifyIntegerObject {
    if g.Err != nil {
        return golifyIntegerObject{
            Value: 0,
            Err:   g.Err,
        }
    }

    i, e := strconv.Atoi(g.Value)
    if e != nil {
        return golifyIntegerObject{
            Value: 0,
            Err: &golifyErr{
                ErrCode: errCode,
                ErrMsg:  errMsg,
            },
        }
    }

    return golifyIntegerObject{
        Value: int64(i),
        Err:   nil,
    }
}
