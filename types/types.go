package types

type PipeConditional struct {
    Branch interface{} `json:"branch,omitempty" yaml:"branch"`
    Tag    interface{} `json:"tag,omitempty"    yaml:"tag"`
}

type PipeJob struct {
    Name  string          `json:"name"  yaml:"name"`
    Image string          `json:"image" yaml:"image"`
    Steps interface{}     `json:"steps" yaml:"steps"`
    When *PipeConditional `json:"when"  yaml:"when"`
}

type PipeConfig struct {
    Jobs []PipeJob `json:"jobs" yaml:"jobs"`
}

type Job struct {
    Name  string   `json:"name"`
    Image string   `json:"image"`
    Steps []string `json:"steps"`
}

type GitInfo struct {
    URL        string `json:"url"`
    RefName    string `json:"ref_name"`
    CommitHash string `json:"commit_hash"`
    IsTag      bool   `json:"is_tag"`
}

type ApiNotification struct {
    PipelineId string  `json:"pipeline_id"`
    Info       GitInfo `json:"git"`
}
