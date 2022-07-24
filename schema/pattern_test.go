package schema_test

import (
    "fmt"
    "regexp"

    "github.com/wolkenwalze/wolkenwalze-plugin-sdk-go/schema"
)

func ExamplePattern() {
    t := schema.Pattern()
    val := "^[a-z]+$"

    var v *regexp.Regexp
    var err error
    v, err = t.Unserialize(val, "val")
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
    // Output: ^[a-z]+$
}
