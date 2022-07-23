package schema

import (
    "fmt"
    "strings"
)

// ErrConstraint indicates that the passed data violated one or more constraints defined in the schema.
type ErrConstraint struct {
    Path    []string
    Message string
    Cause   error
}

// Error creates a printable string for this error.
func (c ErrConstraint) Error() string {
    if len(c.Path) == 0 {
        return fmt.Sprintf("Validation failed: %s", c.Message)
    }
    return fmt.Sprintf("Validation failed for %s: %s", strings.Join(c.Path, " -> "), c.Message)
}

func (c ErrConstraint) Unwrap() error {
    return c.Cause
}

// ErrNoSuchStep indicates that the given step is not supported by a schema.
type ErrNoSuchStep struct {
    Step string
}

// Error creates a printable string for this error.
func (n ErrNoSuchStep) Error() string {
    return fmt.Sprintf("No such step: %s", n.Step)
}

// ErrBadArgument indicates that an invalid configuration was passed to a schema component.
type ErrBadArgument struct {
    Message string
}

// Error creates a printable string for this error.
func (b ErrBadArgument) Error() string {
    return b.Message
}
