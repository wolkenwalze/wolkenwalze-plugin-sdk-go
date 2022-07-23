package schema

type Field[T any] struct {
    Item          Type[T]
    Name          string
    Description   string
    Required      bool
    RequiredIf    []string
    RequiredIfNot []string
    Conflicts     []string
}

type ObjectType[T any] struct {
    Cls        T
    Properties map[string]Field[any]
}

func (o ObjectType[T]) TypeID() TypeID {
    return TypeIDObject
}

func (o ObjectType[T]) Unserialize(data interface{}, path ...string) (T, error) {
    panic("Implement me")
}

func (o ObjectType[T]) Validate(data T, path ...string) error {
    panic("Implement me")
}

func (o ObjectType[T]) Serialize(data T, path ...string) (interface{}, error) {
    panic("Implement me")
}
