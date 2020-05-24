package types

type PipeEngine struct {
    Docker string `json:"docker" yaml:"docker"`
    Vm     string `json:"vm"     yaml:"vm"`
}

type PipeConditional struct {
    Branch interface{} `json:"branch" yaml:"branch"`
    Tag    interface{} `json:"tag"    yaml:"tag"`
}

type PipeJob struct {
    Name   string           `json:"name"   yaml:"name"`
    Steps  interface{}      `json:"steps"  yaml:"steps"`
    Engine *PipeEngine      `json:"engine" yaml:"engine"`
    When   *PipeConditional `json:"when"   yaml:"when"`
}

type PipeConfig struct {
    Jobs []PipeJob `json:"jobs" yaml:"jobs"`
}
