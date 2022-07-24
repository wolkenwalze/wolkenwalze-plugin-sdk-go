package schema

import (
    "fmt"
    "reflect"
)

func List[T any](items Type[T]) ListType[T] {
    return &listType[T]{
        items: items,
    }
}

type ListType[T any] interface {
    Type[[]T]

    Items() Type[T]

    Min() *int
    Max() *int

    WithMin(min int) ListType[T]
    WithMax(max int) ListType[T]
}

type listType[T any] struct {
    items Type[T]
    min   *int
    max   *int
}

func (l listType[T]) Items() Type[T] {
    return l.items
}

func (l listType[T]) Min() *int {
    return l.min
}

func (l listType[T]) Max() *int {
    return l.max
}

func (l *listType[T]) WithMin(min int) ListType[T] {
    l.min = &min
    return l
}

func (l *listType[T]) WithMax(max int) ListType[T] {
    l.max = &max
    return l
}

func (l listType[T]) TypeID() TypeID {
    return TypeIDList
}

func (l listType[T]) Unserialize(data interface{}, path ...string) (result []T, err error) {
    reflectedValue := reflect.ValueOf(data)
    if reflectedValue.Kind() != reflect.Slice {
        return nil, ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("Must be a list, %T given", data),
        }
    }
    result = make([]T, reflectedValue.Len())
    for i := 0; i < reflectedValue.Len(); i++ {
        result[i], err = l.items.Unserialize(
            reflectedValue.Index(i).Interface(),
            append(path, fmt.Sprintf("%d", i))...,
        )
        if err != nil {
            return nil, err
        }
    }
    return result, nil
}

func (l listType[T]) Validate(data []T, path ...string) error {
    for i, v := range data {
        if err := l.items.Validate(v, append(path, fmt.Sprintf("%d", i))...); err != nil {
            return err
        }
    }
    return nil
}

func (l listType[T]) Serialize(data []T, path ...string) (interface{}, error) {
    result := make([]interface{}, len(data))
    for i, v := range data {
        var err error
        result[i], err = l.items.Serialize(v, append(path, fmt.Sprintf("%d", i))...)
        if err != nil {
            return nil, err
        }
    }
    return result, nil
}
