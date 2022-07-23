package schema

import (
    "fmt"
    "reflect"
)

type ListType[T any] struct {
    Items Type[T]
    Min   *int
    Max   *int
}

func (l ListType[T]) TypeID() TypeID {
    return TypeIDList
}

func (l ListType[T]) Unserialize(data interface{}, path ...string) (result []T, err error) {
    reflectedValue := reflect.ValueOf(data)
    if reflectedValue.Kind() != reflect.Slice {
        return nil, ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("Must be a list, %T given", data),
        }
    }
    result = make([]T, reflectedValue.Len())
    for i := 0; i < reflectedValue.Len(); i++ {
        result[i], err = l.Items.Unserialize(
            reflectedValue.Index(i).Interface(),
            append(path, fmt.Sprintf("%d", i))...,
        )
        if err != nil {
            return nil, err
        }
    }
    return result, nil
}

func (l ListType[T]) Validate(data []T, path ...string) error {
    for i, v := range data {
        if err := l.Items.Validate(v, append(path, fmt.Sprintf("%d", i))...); err != nil {
            return err
        }
    }
    return nil
}

func (l ListType[T]) Serialize(data []T, path ...string) (interface{}, error) {
    result := make([]interface{}, len(data))
    for i, v := range data {
        var err error
        result[i], err = l.Items.Serialize(v, append(path, fmt.Sprintf("%d", i))...)
        if err != nil {
            return nil, err
        }
    }
    return result, nil
}
