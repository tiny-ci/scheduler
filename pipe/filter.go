package pipe

import (
    "fmt"
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

func Filter(config *types.PipeConfig, ref string) []types.Job {
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

    return matchedJobs
}
