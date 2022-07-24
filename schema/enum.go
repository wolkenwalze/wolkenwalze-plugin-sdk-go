package schema

import (
    "fmt"
    "reflect"
)

func Enum[T ~string | ~int](
    name string,
    values ...T,
) EnumType[T] {
    rawValues := make(map[T]interface{}, len(values))
    t := reflect.TypeOf(values).Elem()
    var convertFunc func(input T) interface{}
    switch t.Kind() {
    case reflect.String:
        convertFunc = func(input T) interface{} {
            return reflect.ValueOf(input).String()
        }
    case reflect.Int:
        convertFunc = func(input T) interface{} {
            return int(reflect.ValueOf(input).Int())
        }
    default:
        panic(fmt.Errorf("unsupported kind: %v", t.Kind()))
    }
    for _, v := range values {
        rawValues[v] = convertFunc(v)
    }
    return &enumType[T]{
        name:      name,
        values:    values,
        rawValues: rawValues,
    }
}

type EnumType[T ~string | ~int] interface {
    Type[T]

    Name() string
    Values() []T
}

type enumType[T ~string | ~int] struct {
    name      string
    values    []T
    rawValues map[T]interface{}
}

func (e enumType[T]) Name() string {
    return e.name
}

func (e enumType[T]) Values() []T {
    r := make([]T, len(e.values))
    copy(r, e.values)
    return r
}

func (e enumType[T]) TypeID() TypeID {
    return TypeIDEnum
}

func (e enumType[T]) Unserialize(data interface{}, path ...string) (T, error) {
    var defaultValue T
    for k, v := range e.rawValues {
        if v == data || k == data {
            return k, nil
        }
    }
    return defaultValue, ErrConstraint{
        Path:    path,
        Message: fmt.Sprintf("'%v' is not a valid value for the enum '%s'", data, e.name),
    }
}

func (e enumType[T]) Validate(data T, path ...string) error {
    for _, v := range e.values {
        if v == data {
            return nil
        }
    }
    return ErrConstraint{
        Path:    path,
        Message: fmt.Sprintf("'%v' is not a valid value for the enum '%s'", data, e.name),
    }
}

func (e enumType[T]) Serialize(data T, path ...string) (interface{}, error) {
    for k, v := range e.rawValues {
        if k == data {
            return v, nil
        }
    }
    var defaultValue T
    return defaultValue, ErrConstraint{
        Path:    path,
        Message: fmt.Sprintf("'%v' is not a valid value for the enum '%s'", data, e.name),
    }
}
