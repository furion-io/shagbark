package main

import (
	"net/http"
	"time"
)

// Ping targeted URL
func ping(url string) (response, error) {
	// Begin time
	start := time.Now()

	// Send get request
	resp, err := http.Get(url)
	if err != nil {
		return response{}, err
	}
	defer resp.Body.Close()

	// End time
	end := time.Now()

	// Result
	res := response{
		URL:        url,
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Latency:    end.Sub(start).Seconds(),
	}

	return res, nil
}
