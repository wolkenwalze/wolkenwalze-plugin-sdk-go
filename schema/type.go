package schema

import "fmt"

// TypeID is the enum of possible types supported by the protocol.
type TypeID string

const (
    TypeIDEnum    TypeID = "enum"
    TypeIDString  TypeID = "string"
    TypeIDPattern TypeID = "pattern"
    TypeIDInt     TypeID = "integer"
    TypeIDList    TypeID = "list"
    TypeIDMap     TypeID = "map"
    TypeIDObject  TypeID = "object"
)

func (t TypeID) Validate() error {
    switch t {
    case TypeIDEnum:
    case TypeIDString:
    case TypeIDPattern:
    case TypeIDInt:
    case TypeIDList:
    case TypeIDMap:
    case TypeIDObject:
    default:
        return ErrBadArgument{
            fmt.Sprintf("%v is not a valid TypeID", t),
        }
    }
    return nil
}

func (t TypeID) IsMapKey() bool {
    switch t {
    case TypeIDEnum:
        return true
    case TypeIDString:
        return true
    case TypeIDInt:
        return true
    default:
        return false
    }
}

type Type[T any] interface {
    TypeID() TypeID
    Unserialize(data interface{}, path ...string) (T, error)
    Validate(data T, path ...string) error
    Serialize(data T, path ...string) (interface{}, error)
}
