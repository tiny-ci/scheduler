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
