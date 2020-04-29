package main

import (
    "encoding/json"
    "fmt"
    "log"
    "github.com/tiny-ci/core/parser"
    "github.com/tiny-ci/core/pipeconf"
    "github.com/tiny-ci/core/schema"
    "github.com/tiny-ci/core/types"
)

func isStringSlice(intf interface{}) bool {
    switch intf.(type) {
    case []interface{}:
        return true

    default:
        return false
    }
}

func toStringSlice(intf interface{}) []string {
    var ss []string

    if isStringSlice(intf) {
        for _, item := range intf.([]interface{}) {
            ss = append(ss, fmt.Sprintf("%s", item))
        }
    } else {
        ss = append(ss, fmt.Sprintf("%s", intf))
    }

    return ss
}

func appendJob(jobs *[]types.Job, pipeJob *types.PipeJob) {
    *jobs = append(*jobs, types.Job{
        Name:  pipeJob.Name,
        Image: pipeJob.Image,
        Steps: toStringSlice(pipeJob.Steps),
    })
}

func filter(config *types.PipeConfig, ref string) []types.Job {
    matchedJobs := []types.Job{}

    for _, pipeJob := range config.Jobs {
        if pipeJob.When == nil {
            appendJob(&matchedJobs, &pipeJob)
            continue
        }

        if pipeJob.When.Branch == nil && pipeJob.When.Tag == nil {
            continue
        }

        var conditionals []string

        if pipeJob.When.Branch != nil {
            conditionals = append(conditionals, toStringSlice(pipeJob.When.Branch)...)
        }

        if pipeJob.When.Tag != nil {
            conditionals = append(conditionals, toStringSlice(pipeJob.When.Tag)...)
        }

        hasMatch := false
        for _, cond := range conditionals {
            if ref == cond {
                appendJob(&matchedJobs, &pipeJob)
                hasMatch = true
                break
            }
        }

        if hasMatch { continue }
    }

    fmt.Println(matchedJobs)

    return matchedJobs
}

func main() {
    configContent, err := pipeconf.Fetch(&pipeconf.GitRef{
        Name: "master",
        URL:  "https://github.com/tiny-ci/example",
        Hash: "961240b68021e7d8ecb53c0b5b4e1e8e097fb319",
        IsTag: false,
    })

    if err != nil {
        log.Fatal(err)
    }

    config, err := parser.ParsePipeConfig(configContent.Bytes())
    if err != nil {
        log.Println("parser error")
        log.Fatal(err)
    }

    var marshalledConfig []byte
    marshalledConfig, err = json.Marshal(config)
    if err != nil {
        log.Println("marshall error")
        log.Fatal(err)
    }

    err = schema.ValidateConfig(marshalledConfig)
    if err != nil {
        log.Println("validation error")
        log.Fatal(err)
    }

    filter(config, "master")
}
