package schema

import (
    "fmt"
    "reflect"
    "strings"
)

func Field[T any](id string, item Type[T]) FieldType {
    return &fieldType[T]{
        id:   id,
        name: id,
        item: item,
    }
}

// FieldType describes a field inside an object. It is not a generic type to enable fields of different types being
// added to an object, which Go only supports without a generic type.
type FieldType interface {
    ID() string
    Name() string
    Description() string
    Required() bool
    RequiredIf() []string
    RequiredIfNot() []string
    Conflicts() []string
    Item() interface{}

    // TypeID returns the underlying type for this field.
    TypeID() TypeID
    Unserialize(data interface{}, path ...string) (any, error)
    Validate(data any, path ...string) error
    Serialize(data any, path ...string) (interface{}, error)
}

type fieldType[T any] struct {
    id            string
    name          string
    description   string
    item          Type[T]
    required      bool
    requiredIf    []string
    requiredIfNot []string
    conflicts     []string
}

func (f fieldType[T]) ID() string {
    return f.id
}

func (f fieldType[T]) Name() string {
    return f.name
}

func (f fieldType[T]) Description() string {
    return f.description
}

func (f fieldType[T]) Required() bool {
    return f.required
}

func (f fieldType[T]) RequiredIf() []string {
    return f.requiredIf
}

func (f fieldType[T]) RequiredIfNot() []string {
    return f.requiredIfNot
}

func (f fieldType[T]) Conflicts() []string {
    return f.conflicts
}

func (f fieldType[T]) Item() interface{} {
    return f.item
}

func (f fieldType[T]) TypeID() TypeID {
    return f.item.TypeID()
}

func (f fieldType[T]) Unserialize(data interface{}, path ...string) (any, error) {
    return f.item.Unserialize(data, path...)
}

func (f fieldType[T]) Validate(data any, path ...string) error {
    return f.item.Validate(data.(T), path...)
}

func (f fieldType[T]) Serialize(data any, path ...string) (interface{}, error) {
    return f.item.Serialize(data.(T), path...)
}

func Object[T any](properties ...FieldType) ObjectType[T] {
    fields := make(map[string]FieldType, len(properties))
    validKeys := make([]string, len(properties))
    for i, field := range properties {
        if _, ok := fields[field.ID()]; ok {
            panic(ErrBadArgument{
                fmt.Sprintf("duplicate field: %s", field.ID()),
            })
        }
        fields[field.ID()] = field
        validKeys[i] = field.ID()
    }

    return &objectType[T]{
        validKeys:  validKeys,
        properties: fields,
    }
}

type ObjectType[T any] interface {
    Type[T]

    Properties() map[string]FieldType
}

type objectType[T any] struct {
    validKeys  []string
    properties map[string]FieldType
}

func (o objectType[T]) TypeID() TypeID {
    return TypeIDObject
}

func (o objectType[T]) Unserialize(data interface{}, path ...string) (result T, err error) {
    dataType := reflect.TypeOf(data)
    dataValue := reflect.ValueOf(data)
    if dataType.Kind() != reflect.Map {
        return result, ErrConstraint{
            Path:    path,
            Message: fmt.Sprintf("must be a map, got %T", data),
        }
    }
    resultValue := reflect.ValueOf(&result)
    for _, key := range dataValue.MapKeys() {
        value := dataValue.MapIndex(key)
        if key.Kind() != reflect.String {
            return result, ErrConstraint{
                Path:    path,
                Message: fmt.Sprintf(""),
            }
        }

        keyName := key.String()
        field, ok := o.properties[keyName]
        if !ok {
            return result, ErrConstraint{
                Path: append(path, keyName),
                Message: fmt.Sprintf(
                    "invalid parameter '%s', expected one of: '%s'",
                    keyName,
                    strings.Join(o.validKeys, "', '"),
                ),
            }
        }
        unserializedValue, err := field.Unserialize(value.Interface(), append(path, keyName)...)
        if err != nil {
            return result, err
        }
        resultValue.Elem().FieldByName(keyName).Set(reflect.ValueOf(unserializedValue))
    }
    return result, o.Validate(result, path...)
}

func (o objectType[T]) Validate(data T, path ...string) error {
    dataValue := reflect.ValueOf(data)
    for key, field := range o.properties {
        val := dataValue.FieldByName(key).Interface()
        path := append(path, key)
        if val == nil {
            if field.Required() {
                return ErrConstraint{
                    Path:    path,
                    Message: "field is required but not set",
                }
            }
            for _, r := range field.RequiredIf() {
                if dataValue.FieldByName(r).Interface() != nil {
                    return ErrConstraint{
                        Path:    path,
                        Message: fmt.Sprintf("field is required because '%s' is set", r),
                    }
                }
            }
            if len(field.RequiredIfNot()) > 0 {
                noneSet := true
                for _, r := range field.RequiredIfNot() {
                    if dataValue.FieldByName(r).Interface() != nil {
                        noneSet = false
                        break
                    }
                }
                if noneSet {
                    return ErrConstraint{
                        Path: path,
                        Message: fmt.Sprintf(
                            "field is required because none of '%s' is set",
                            strings.Join(field.RequiredIfNot(), "', '"),
                        ),
                    }
                }
            }
        } else {
            for _, c := range field.Conflicts() {
                if dataValue.FieldByName(c).Interface() != nil {
                    return ErrConstraint{
                        Path: path,
                        Message: fmt.Sprintf(
                            "field conflicts '%s', set one of the two, not both",
                            c,
                        ),
                    }
                }
            }
            if err := field.Validate(val, path...); err != nil {
                return err
            }
        }
    }
    return nil
}

func (o objectType[T]) Serialize(data T, path ...string) (interface{}, error) {
    if err := o.Validate(data, path...); err != nil {
        return nil, err
    }
    result := map[string]interface{}{}
    dataValue := reflect.ValueOf(data)
    for key, field := range o.properties {
        val := dataValue.FieldByName(key).Interface()
        if val == nil {
            continue
        }
        serializedValue, err := field.Serialize(val, append(path, key)...)
        if err != nil {
            return nil, err
        }
        result[key] = serializedValue
    }
    return result, nil
}

func (o objectType[T]) Properties() map[string]FieldType {
    return o.properties
}
