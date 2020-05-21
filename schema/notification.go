package schema

var notificationSchema = []byte(`{
    "type": "object",
    "required": ["pipeline_id", "git"],
    "additionalProperties": false,
    "properties": {
        "pipeline_id": { "type": "string" },
        "git": {
            "type": "object",
            "required": ["url", "ref_name", "commit_hash", "is_tag"],
            "additionalProperties": false,
            "properties": {
                "url": { "type": "string" },
                "ref_name": { "type": "string" },
                "commit_hash": { "type": "string" },
                "is_tag": { "type": "boolean" }
            }
        }
    }}`)


func ValidateNotification(data []byte) error {
    return validate(notificationSchema, data)
}
