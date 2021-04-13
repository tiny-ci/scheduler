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
                        "oneOf": [ { "required": ["docker"] }, { "required": ["vm"] } ],
                        "additionalProperties": false,
                        "properties": {
                            "docker": { "type": "string" },
                            "vm": { "type": "string" }
                        }
                    },
                    "steps": {
                        "oneOf": [
                            { "type": "string" },
                            { "type": "array", "items": { "type": "string" } }
                        ]
                    },
                    "when": {
                        "oneOf": [
                            { "type": "null" },
                            {
                                "allOf": [
                                    { "type": "object" },
                                    { "oneOf": [ { "required": ["tag"] }, { "required": ["branch"] } ] }
                                ]
                            }
                        ],
                        "additionalProperties": false,
                        "properties": {
                            "tag": {
                                "oneOf": [
                                    { "type": "string" },
                                    { "type": "array", "items": { "type": "string" } }
                                ]
                            },
                            "branch": {
                                "oneOf": [
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
