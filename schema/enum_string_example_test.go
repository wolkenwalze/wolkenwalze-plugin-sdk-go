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

func ExampleEnumType_string() {
    enumType := schema.Enum("Color", ColorGreen, ColorRed)

    // Unserialize a string value to the enum
    color, err := enumType.Unserialize("green")
    if err != nil {
        panic(err)
    }
    if color != ColorGreen {
        panic(fmt.Errorf("incorrect enum value returned: %s", color))
    }

    // Validate unserialized value
    if err := enumType.Validate(color); err != nil {
        panic(err)
    }

    // Serialize value to a raw type
    serializedValue, err := enumType.Serialize(color)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%v", serializedValue)
    // Output: green
}
