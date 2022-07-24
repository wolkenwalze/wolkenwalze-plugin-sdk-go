package schema

import (
    "fmt"
    "reflect"
)

func Map[K MapKey, V any](keys Type[K], values Type[V]) MapType[K, V] {
    return &mapType[K, V]{
        keys:   keys,
        values: values,
    }
}

type MapKey interface {
    int | string
}

type MapType[K MapKey, V any] interface {
    Type[map[K]V]

    Keys() Type[K]
    Values() Type[V]

    Min() *int
    Max() *int

    WithMin(min int) MapType[K, V]
    WithMax(max int) MapType[K, V]
}

type mapType[K MapKey, V any] struct {
    keys   Type[K]
    values Type[V]
    min    *int
    max    *int
}

func (m mapType[K, V]) Keys() Type[K] {
    return m.keys
}

func (m mapType[K, V]) Values() Type[V] {
    return m.values
}

func (m mapType[K, V]) Min() *int {
    return m.min
}

func (m mapType[K, V]) Max() *int {
    return m.max
}

func (m *mapType[K, V]) WithMin(min int) MapType[K, V] {
    m.min = &min
    return m
}

func (m *mapType[K, V]) WithMax(max int) MapType[K, V] {
    m.max = &max
    return m
}

func (m mapType[K, V]) TypeID() TypeID {
    return TypeIDMap
}

func (m mapType[K, V]) Unserialize(data interface{}, path ...string) (result map[K]V, err error) {
    reflectedValue := reflect.ValueOf(data)
    if reflectedValue.Kind() != reflect.Map {
        return nil, ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("must be a map, %T given", data),
        }
    }
    result = make(map[K]V, reflectedValue.Len())
    for _, key := range reflectedValue.MapKeys() {
        k := key.Interface()
        v := reflectedValue.MapIndex(key).Interface()
        newPath := append(path, fmt.Sprintf("%v", k))
        key, err := m.keys.Unserialize(k, append(newPath, "key")...)
        if err != nil {
            return nil, err
        }
        value, err := m.values.Unserialize(v, append(newPath, "value")...)
        if err != nil {
            return nil, err
        }
        result[key] = value
    }
    return result, nil
}

func (m mapType[K, V]) Validate(data map[K]V, path ...string) error {
    for k, v := range data {
        newPath := append(path, fmt.Sprintf("%v", k))

        if err := m.keys.Validate(k, append(newPath, "key")...); err != nil {
            return err
        }
        if err := m.values.Validate(v, append(newPath, "value")...); err != nil {
            return err
        }
    }
    return nil
}

func (m mapType[K, V]) Serialize(data map[K]V, path ...string) (interface{}, error) {
    result := make(map[interface{}]interface{}, len(data))
    for k, v := range data {
        newPath := append(path, fmt.Sprintf("%v", k))
        key, err := m.keys.Serialize(k, append(newPath, "key")...)
        if err != nil {
            return nil, err
        }
        value, err := m.values.Serialize(v, append(newPath, "value")...)
        if err != nil {
            return nil, err
        }
        result[key] = value
    }
    return result, nil
}
