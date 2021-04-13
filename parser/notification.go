package parser

import (
	"encoding/json"
	"github.com/tiny-ci/core/types"
)

func ParseNotification(content []byte) (*types.ApiNotification, error) {
	var notification types.ApiNotification

	err := json.Unmarshal(content, &notification)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}
