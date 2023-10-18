package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var endpoint = "/time_entries.json"

type RedmineClient struct {
	host   string
	apiKey string
}

type RequestBody struct {
	TimeEntry TimeEntry `json:"time_entry"`
}

var ContentError = errors.New("Content Error")

func (r *RedmineClient) saveTimeEntry(timeEntry TimeEntry) error {
	url := r.host + endpoint

	v, err := json.Marshal(RequestBody{TimeEntry: timeEntry})
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(v))
	if err != nil {
		return fmt.Errorf("failed to create a request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Redmine-API-Key", r.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnprocessableEntity {
		return ContentError
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to request: http status code %d", resp.StatusCode)
	}

	return nil
}
