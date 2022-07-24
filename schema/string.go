package schema

import (
    "fmt"
    "regexp"
)

func String() StringType {
    return &stringType{}
}

type StringType interface {
    Type[string]

    MinLength() *int
    MaxLength() *int
    Pattern() *regexp.Regexp

    WithMinLength(int) StringType
    WithMaxLength(int) StringType
    WithPattern(regexp2 *regexp.Regexp) StringType
}

type stringType struct {
    minLength *int
    maxLength *int
    pattern   *regexp.Regexp
}

func (s *stringType) MinLength() *int {
    return s.minLength
}

func (s *stringType) MaxLength() *int {
    return s.maxLength
}

func (s *stringType) Pattern() *regexp.Regexp {
    return s.pattern
}

func (s *stringType) WithMinLength(min int) StringType {
    s.minLength = &min
    return s
}

func (s *stringType) WithMaxLength(max int) StringType {
    s.maxLength = &max
    return s
}

func (s *stringType) WithPattern(re *regexp.Regexp) StringType {
    if re == nil {
        panic(fmt.Errorf("nil parameter passed to StringType.WithPattern()"))
    }
    s.pattern = re
    return s
}

func (s stringType) TypeID() TypeID {
    return TypeIDString
}

func (s stringType) Unserialize(data interface{}, path ...string) (result string, err error) {
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

func (s stringType) Validate(data string, path ...string) error {
    if s.minLength != nil && len(data) < *s.minLength {
        return ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("string must be at least %d characters, %d given", *s.minLength, len(data)),
        }
    }
    if s.maxLength != nil && len(data) > *s.maxLength {
        return ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("string must be at most %d characters, %d given", *s.maxLength, len(data)),
        }
    }
    if s.pattern != nil && !s.pattern.MatchString(data) {
        return ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("string must match pattern %s", s.pattern.String()),
        }
    }
    return nil
}

func (s stringType) Serialize(data string, path ...string) (interface{}, error) {
    return data, s.Validate(data, path...)
}
