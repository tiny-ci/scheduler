package types

type Job struct {
    Name   string   `json:"name"`
    Engine string   `json:"engine"`
    Image  string   `json:"image"`
    Steps  []string `json:"steps"`
}
