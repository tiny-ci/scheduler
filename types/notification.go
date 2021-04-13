package types

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
