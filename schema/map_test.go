package schema_test

import (
    "fmt"

    "github.com/wolkenwalze/wolkenwalze-plugin-sdk-go/schema"
)

func ExampleMapType() {
    val := map[string]int{
        "test": 5,
    }
    t := schema.MapType[string, int]{
        Keys:   schema.StringType{},
        Values: schema.IntType{},
    }

    v, err := t.Unserialize(val, "val")
    if err != nil {
        panic(err)
    }

    if err := t.Validate(v, "val"); err != nil {
        panic(err)
    }

    serialized, err := t.Serialize(v, "val")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%v", serialized)
    // Output: map[test:5]
}
