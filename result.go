package main

import (
	"encoding/json"
)

// response define ping result information
type response struct {
	URL        string  `json:"url"`
	Status     string  `json:"status"`
	StatusCode int     `json:"status_code"`
	Latency    float64 `json:"latency"`
}

// toJSON convert response into valid JSON
func (r *response) toJSON() []byte {
	j, _ := json.Marshal(r)
	return j
}
