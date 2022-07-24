package schema_test

import (
    "fmt"

    "github.com/wolkenwalze/wolkenwalze-plugin-sdk-go/schema"
)

type ByteMultiplier int

const (
    ByteMultiplierKB ByteMultiplier = 1024
    ByteMultiplierMB ByteMultiplier = 1048576
)

func ExampleEnumType_int() {
    enumType := schema.Enum("ByteMultiplier", ByteMultiplierKB, ByteMultiplierMB)
    val := 1024

    // Unserialize a string value to the enum
    byteMultiplier, err := enumType.Unserialize(val, "val")
    if err != nil {
        panic(err)
    }
    if byteMultiplier != ByteMultiplierKB {
        panic(fmt.Errorf("incorrect enum value returned: %d", byteMultiplier))
    }

    // Validate unserialized value
    if err := enumType.Validate(byteMultiplier, "val"); err != nil {
        panic(err)
    }

    // Serialize value to a raw type
    serializedValue, err := enumType.Serialize(byteMultiplier, "val")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%v", serializedValue)
    // Output: 1024
}
