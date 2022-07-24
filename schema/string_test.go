package schema_test

import (
    "fmt"
    "regexp"

    "github.com/wolkenwalze/wolkenwalze-plugin-sdk-go/schema"
)

func ExampleStringType() {
    t := schema.String()
    val := "Hello world!"
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
    // Output: Hello world!
}

func ExampleStringType_min() {
    t := schema.String().WithMinLength(5)
    val := "foo"
    _, err := t.Unserialize(val, "val")
    if err != nil {
        fmt.Printf("Invalid input: %v", err)
    }
    // Output: Invalid input: Validation failed for val: string must be at least 5 characters, 3 given
}

func ExampleStringType_max() {
    t := schema.String().WithMaxLength(5)
    val := "Hello world!"
    _, err := t.Unserialize(val, "val")
    if err != nil {
        fmt.Printf("Invalid input: %v", err)
    }
    // Output: Invalid input: Validation failed for val: string must be at most 5 characters, 12 given
}

func ExampleStringType_pattern() {
    t := schema.String().WithPattern(regexp.MustCompile(`^[0-9]+$`))
    val := "Hello world!"
    _, err := t.Unserialize(val, "val")
    if err != nil {
        fmt.Printf("Invalid input: %v", err)
    }
    // Output: Invalid input: Validation failed for val: string must match pattern ^[0-9]+$
}
