package schema

import (
    "fmt"
    "strconv"
)

func Int() IntType {
    return &intType{}
}

type IntType interface {
    Type[int]

    Min() *int
    Max() *int

    WithMin(min int) IntType
    WithMax(max int) IntType
}

type intType struct {
    min *int
    max *int
}

func (i intType) Min() *int {
    return i.min
}

func (i intType) Max() *int {
    return i.max
}

func (i *intType) WithMin(min int) IntType {
    i.min = &min
    return i
}

func (i *intType) WithMax(max int) IntType {
    i.max = &max
    return i
}

func (i intType) TypeID() TypeID {
    return TypeIDInt
}

func (i intType) Unserialize(data interface{}, path ...string) (result int, err error) {
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

func (i intType) Validate(data int, path ...string) error {
    if i.min != nil && data < *i.min {
        return ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("must be at least %d", *i.min),
        }
    }
    if i.max != nil && data > *i.max {
        return ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("must be at most %d", *i.min),
        }
    }
    return nil
}

func (i intType) Serialize(data int, path ...string) (interface{}, error) {
    return data, i.Validate(data)
}
