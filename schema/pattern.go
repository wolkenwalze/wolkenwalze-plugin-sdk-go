package schema

import (
    "fmt"
    "regexp"
)

func Pattern() PatternType {
    return &patternType{}
}

type PatternType interface {
    Type[*regexp.Regexp]
}

type patternType struct {
}

func (p patternType) TypeID() TypeID {
    return TypeIDPattern
}

func (p patternType) Unserialize(data interface{}, path ...string) (*regexp.Regexp, error) {
    switch d := data.(type) {
    case string:
        re, err := regexp.Compile(d)
        if err != nil {
            return nil, ErrConstraint{
                Path:    path,
                Message: fmt.Sprintf("invalid regular expression (%v)", err),
                Cause:   err,
            }
        }
        return re, nil
    default:
        return nil, ErrConstraint{
            Path:    path,
            Message: "must be a string",
        }
    }
}

func (p patternType) Validate(data *regexp.Regexp, path ...string) error {
    if data == nil {
        return ErrConstraint{
            Path:    path,
            Message: "must not be nil",
        }
    }
    return nil
}

func (p patternType) Serialize(data *regexp.Regexp, path ...string) (interface{}, error) {
    if data == nil {
        return nil, ErrConstraint{
            Path:    path,
            Message: "must not be nil",
        }
    }
    return data.String(), nil
}
