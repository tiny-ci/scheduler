package pipe

import (
	"fmt"
	"github.com/tiny-ci/core/types"
	"strings"
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
	engine := "docker"
	image := "ubuntu"

	if pipeJob.Engine != nil {
		if pipeJob.Engine.Vm != "" {
			engine = "vm"
			image = pipeJob.Engine.Vm
		} else {
			image = pipeJob.Engine.Docker
		}
	}

	*jobs = append(*jobs, types.Job{
		Name:   pipeJob.Name,
		Engine: engine,
		Image:  image,
		Steps:  toStringSlice(pipeJob.Steps),
	})
}

func Filter(config *types.PipeConfig, ref string) ([]types.Job, error) {
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

		for _, cond := range conditionals {
			if ref == cond {
				appendJob(&matchedJobs, &pipeJob)
				break
			}

			if strings.HasPrefix(cond, "\\") {
				hasMatch, err := EvalExpr(cond, ref)
				if err != nil {
					return nil, err
				}

				if hasMatch {
					appendJob(&matchedJobs, &pipeJob)
					break
				}
			}
		}
	}

	return matchedJobs, nil
}
