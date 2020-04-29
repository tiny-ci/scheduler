package schema

import (
    "errors"
    "encoding/json"
    "log"
    "github.com/qri-io/jsonschema"
)

var configSchema = []byte(`{
    "type": "object",
    "required": ["jobs"],
    "additionalProperties": false,
    "properties": {
        "jobs": {
            "type": "array",
            "minItems": 1,
            "maxItems": 10,
            "items": {
                "type": "object",
                "required": ["name", "image", "steps"],
                "additionalProperties": false,
                "properties": {
                    "name": { "type": "string" },
                    "image": { "type": "string" },

                    "steps": {
                        "anyOf": [
                            { "type": "string" },
                            { "type": "array", "items": { "type": "string" } }
                        ]
                    },
                    "when": {
                        "anyOf": [ { "type": "null" }, { "type": "object" } ],
                        "anyOf": [ { "required": ["tag"] }, { "required": ["branch"] } ],
                        "properties": {
                            "tag": {
                                "anyOf": [
                                    { "type": "string" },
                                    { "type": "array", "items": { "type": "string" } }
                                ]
                            },
                            "branch": {
                                "anyOf": [
                                    { "type": "string" },
                                    { "type": "array", "items": { "type": "string" } }
                                ]
                            }
                        }
                    }

                }
            }
        }
    }}`)

func validate(schema []byte, data []byte) error {
    rs := &jsonschema.RootSchema{}
    if err := json.Unmarshal(schema, rs); err != nil {
        return err
    }

    if jsErrors, _ := rs.ValidateBytes(data); len(jsErrors) > 0 {
        for _, err := range jsErrors {
            log.Println(err)
        }

        return errors.New("")
    }

    return nil
}

func ValidateConfig(data []byte) error {
    return validate(configSchema, data)
}
