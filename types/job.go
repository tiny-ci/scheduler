package types

type Job struct {
    Name  string   `json:"name"`
    Image string   `json:"image"`
    Steps []string `json:"steps"`
}
