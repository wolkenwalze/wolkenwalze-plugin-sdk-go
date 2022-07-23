package schema_test

import (
    "fmt"

    "github.com/wolkenwalze/wolkenwalze-plugin-sdk-go/schema"
)

func ExampleListType() {
    t := schema.ListType[string]{
        Items: schema.StringType{},
    }

    var val interface{}
    val = []string{
        "Hello world!",
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
    // Output: [Hello world!]
}
