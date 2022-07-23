package schema

import (
    "fmt"
    "reflect"
)

type MapKey interface {
    int | string
}

type MapType[K MapKey, V any] struct {
    Keys   Type[K]
    Values Type[V]
    Min    *int
    Max    *int
}

func (m MapType[K, V]) TypeID() TypeID {
    return TypeIDMap
}

func (m MapType[K, V]) Unserialize(data interface{}, path ...string) (result map[K]V, err error) {
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
        key, err := m.Keys.Unserialize(k, append(newPath, "key")...)
        if err != nil {
            return nil, err
        }
        value, err := m.Values.Unserialize(v, append(newPath, "value")...)
        if err != nil {
            return nil, err
        }
        result[key] = value
    }
    return result, nil
}

func (m MapType[K, V]) Validate(data map[K]V, path ...string) error {
    for k, v := range data {
        newPath := append(path, fmt.Sprintf("%v", k))

        if err := m.Keys.Validate(k, append(newPath, "key")...); err != nil {
            return err
        }
        if err := m.Values.Validate(v, append(newPath, "value")...); err != nil {
            return err
        }
    }
    return nil
}

func (m MapType[K, V]) Serialize(data map[K]V, path ...string) (interface{}, error) {
    result := make(map[interface{}]interface{}, len(data))
    for k, v := range data {
        newPath := append(path, fmt.Sprintf("%v", k))
        key, err := m.Keys.Serialize(k, append(newPath, "key")...)
        if err != nil {
            return nil, err
        }
        value, err := m.Values.Serialize(v, append(newPath, "value")...)
        if err != nil {
            return nil, err
        }
        result[key] = value
    }
    return result, nil
}
