package schema

import (
    "fmt"
    "strconv"
)

type IntType struct {
    Min *int
    Max *int
}

func (i IntType) TypeID() TypeID {
    return TypeIDInt
}

func (i IntType) Unserialize(data interface{}, path ...string) (result int, err error) {
    switch d := data.(type) {
    case string:
        result, err = strconv.Atoi(d)
        if err != nil {
            return 0, ErrConstraint{
                Path:    path,
                Message: "Must be an integer",
                Cause:   err,
            }
        }
    case int:
        result = d
    case int8:
        result = int(d)
    case int16:
        result = int(d)
    case int32:
        result = int(d)
    case int64:
        result = int(d)
    case uint8:
        result = int(d)
    case uint16:
        result = int(d)
    case uint32:
        result = int(d)
    case uint64:
        result = int(d)
    default:
        return 0, ErrConstraint{
            Path:    path,
            Message: "Must be an integer",
        }
    }
    return result, i.Validate(result, path...)
}

func (i IntType) Validate(data int, path ...string) error {
    if i.Min != nil && data < *i.Min {
        return ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("must be at least %d", *i.Min),
        }
    }
    if i.Max != nil && data > *i.Max {
        return ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("must be at most %d", *i.Min),
        }
    }
    return nil
}

func (i IntType) Serialize(data int, path ...string) (interface{}, error) {
    return data, i.Validate(data)
}
