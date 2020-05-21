package parser

import (
    "encoding/json"
    "github.com/tiny-ci/core/types"
    "gopkg.in/yaml.v2"
)

func ParsePipeConfig(content []byte) (*types.PipeConfig, error) {
    config := types.PipeConfig{}

    err := yaml.Unmarshal(content, &config)
    if err != nil {
        return nil, err
    }

    return &config, nil
}

func ParseNotification(content []byte) (*types.ApiNotification, error) {
    var notification types.ApiNotification

    err := json.Unmarshal(content, &notification)
    if err != nil {
        return nil, err
    }

    return &notification, nil
}
