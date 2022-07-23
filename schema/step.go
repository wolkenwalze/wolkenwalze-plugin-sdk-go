package schema

func NewStep[T any](
    id string,
    name string,
    description string,
    results map[string]interface{},
    handler func(params T) (string, interface{}),
) Step {
    panic("Implement me")
}

type Step interface {
    ID() string
    Name() string
    Description() string
    Execute(params interface{}) (string, interface{})
}
