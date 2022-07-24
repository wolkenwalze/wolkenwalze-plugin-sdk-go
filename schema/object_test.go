package schema_test

import (
    "fmt"

    "github.com/wolkenwalze/wolkenwalze-plugin-sdk-go/schema"
)

type TestObject struct {
    A string
    B int
    C []string
    D map[string]int
}

func ExampleObject() {
    t := schema.Object[TestObject](
        schema.Field[string](
            "A",
            schema.String(),
        ),
        schema.Field[int](
            "B",
            schema.Int(),
        ),
        schema.Field[[]string](
            "C",
            schema.List[string](schema.String()),
        ),
        schema.Field[map[string]int](
            "D",
            schema.Map[string, int](schema.String(), schema.Int()),
        ),
    )
    data := map[string]interface{}{
        "A": "Hello world!",
        "B": 42,
        "C": []string{"foo"},
        "D": map[string]int{
            "bar": 42,
        },
    }
    unserializedData, err := t.Unserialize(data, "data")
    if err != nil {
        panic(err)
    }

    if err := t.Validate(unserializedData, "data"); err != nil {
        panic(err)
    }

    serializedData, err := t.Serialize(unserializedData, "data")
    if err != nil {
        panic(err)
    }

    if serializedData.(map[string]interface{})["A"] != "Hello world!" {
        panic(fmt.Errorf("invalid A value returned"))
    }
    // Output:
}
