package schema_test

import (
    "fmt"

    "github.com/wolkenwalze/wolkenwalze-plugin-sdk-go/schema"
)

func ExampleIntType() {
    t := schema.Int()
    val := 42
    v, err := t.Unserialize(val, "val")
    if err != nil {
        panic(err)
    }

    if err := t.Validate(v); err != nil {
        panic(err)
    }

    serialized, err := t.Serialize(v, "val")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%v", serialized)
    // Output: 42
}
