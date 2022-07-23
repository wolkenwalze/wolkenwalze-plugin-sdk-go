package main

import (
    "fmt"
    "os"
    "regexp"

    "github.com/wolkenwalze/wolkenwalze-plugin-sdk-go/plugin"
    "github.com/wolkenwalze/wolkenwalze-plugin-sdk-go/schema"
)

type PodScenarioParams struct {
    NamespacePattern regexp.Regexp `default:".*"`
    PodNamePattern   regexp.Regexp `default:".*"`
}

type Pod struct {
    Namespace string
    Name      string
}

type PodScenarioResults struct {
    PodsKilled []Pod
}

type PodScenarioError struct {
    Error string
}

func PodScenario(params PodScenarioParams) (string, interface{}) {
    return "error", PodScenarioError{
        fmt.Sprintf(
            "Cannot kill pod %s in namespace %s, function not implemented",
            params.PodNamePattern.String(),
            params.NamespacePattern.String(),
        ),
    }
}

func main() {
    os.Exit(plugin.Run(plugin.BuildSchema(
        schema.NewStep(
            "pod",
            "Pod scenario",
            "Kill one or more pods matching the criteria",
            map[string]interface{}{
                "success": PodScenarioResults{},
                "error":   PodScenarioError{},
            },
            PodScenario,
        ),
    )))
}
