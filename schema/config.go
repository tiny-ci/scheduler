package schema

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
                "required": ["name", "steps"],
                "additionalProperties": false,
                "properties": {
                    "name": { "type": "string" },
                    "engine": {
                        "type": "object",
                        "oneOf": [ { "required": "docker" }, { "required": "vm" } ],
                        "properties": {
                            "docker": { "type": "string" },
                            "vm": { "type": "string" }
                        }
                    },
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

func ValidateConfig(data []byte) error {
    return validate(configSchema, data)
}
