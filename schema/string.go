package schema

import (
    "fmt"
    "regexp"
)

type StringType struct {
    MinLength *int
    MaxLength *int
    Pattern   *regexp.Regexp
}

func (s StringType) TypeID() TypeID {
    return TypeIDString
}

func (s StringType) Unserialize(data interface{}, path ...string) (result string, err error) {
    switch d := data.(type) {
    case string:
        result = d
    case int:
        result = fmt.Sprintf("%d", d)
    case int8:
        result = fmt.Sprintf("%d", d)
    case int16:
        result = fmt.Sprintf("%d", d)
    case int32:
        result = fmt.Sprintf("%d", d)
    case int64:
        result = fmt.Sprintf("%d", d)
    case uint8:
        result = fmt.Sprintf("%d", d)
    case uint16:
        result = fmt.Sprintf("%d", d)
    case uint32:
        result = fmt.Sprintf("%d", d)
    case uint64:
        result = fmt.Sprintf("%d", d)
    default:
        return "", ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("must be a string, %T given", data),
        }
    }
    return result, s.Validate(result, path...)
}

func (s StringType) Validate(data string, path ...string) error {
    if s.MinLength != nil && len(data) < *s.MinLength {
        return ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("string must be at least %d characters, %d given", *s.MinLength, len(data)),
        }
    }
    if s.MaxLength != nil && len(data) > *s.MaxLength {
        return ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("string must be at most %d characters, %d given", *s.MinLength, len(data)),
        }
    }
    if s.Pattern != nil && !s.Pattern.MatchString(data) {
        return ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("string must match pattern %s", s.Pattern.String()),
        }
    }
    return nil
}

func (s StringType) Serialize(data string, path ...string) (interface{}, error) {
    return data, s.Validate(data, path...)
}
