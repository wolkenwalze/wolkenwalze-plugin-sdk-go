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

func ExampleNewIntEnum() {
    enumType := schema.NewIntEnum("ByteMultiplier", ByteMultiplierKB, ByteMultiplierMB)
    val := 1024
    byteMultiplier, err := enumType.Unserialize(val, "val")
    if err != nil {
        panic(err)
    }

    if byteMultiplier != ByteMultiplierKB {
        panic(fmt.Errorf("incorrect enum value returned: %d", byteMultiplier))
    }

    if err := enumType.Validate(byteMultiplier, "val"); err != nil {
        panic(err)
    }

    serializedValue, err := enumType.Serialize(byteMultiplier, "val")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%v", serializedValue)
    // Output: 1024
}
