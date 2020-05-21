package parser

import (
    "gopkg.in/yaml.v2"
    "github.com/tiny-ci/core/types"
)

func ParsePipeConfig(content []byte) (*types.PipeConfig, error) {
    config := types.PipeConfig{}

    err := yaml.Unmarshal(content, &config)
    if err != nil {
        return nil, err
    }

    return &config, nil
}
