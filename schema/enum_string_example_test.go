package schema_test

import (
    "fmt"

    "github.com/wolkenwalze/wolkenwalze-plugin-sdk-go/schema"
)

type Color string

const (
    ColorGreen Color = "green"
    ColorRed   Color = "red"
)

func ExampleNewStringEnum() {
    enumType := schema.NewStringEnum("Color", ColorGreen, ColorRed)
    color, err := enumType.Unserialize("green")
    if err != nil {
        panic(err)
    }
    if color != ColorGreen {
        panic(fmt.Errorf("incorrect enum value returned: %s", color))
    }

    if err := enumType.Validate(color); err != nil {
        panic(err)
    }

    serializedValue, err := enumType.Serialize(color)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%v", serializedValue)
    // Output: green
}
