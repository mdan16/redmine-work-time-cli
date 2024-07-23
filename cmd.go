package main

import (
	"fmt"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
	"gopkg.in/yaml.v3"
)

type CMD struct {
	outStream, errStream io.Writer
	input                Input
}

type Input struct {
	APIHost       string      `yaml:"apiHost"`
	RedmineAPIKey string      `yaml:"redmineApiKey"`
	SpentOn       string      `yaml:"spentOn"`
	TimeEntries   []TimeEntry `yaml:"timeEntries"`
}

type TimeEntry struct {
	IssueId    int     `yaml:"issueId" json:"issue_id"`
	SpentOn    string  `json:"spent_on"`
	Hours      float32 `yaml:"hours" json:"hours"`
	ActivityId int     `yaml:"activityId" json:"activity_id"`
	Comments   string  `yaml:"comments" json:"comments"`
}

func (c *CMD) Run(file string) error {
	if err := c.parseInput(file); err != nil {
		return err
	}

	client := RedmineClient{
		host:   c.input.APIHost,
		apiKey: c.input.RedmineAPIKey,
	}

	bar := progressbar.Default(int64(len(c.input.TimeEntries)))
	for _, t := range c.input.TimeEntries {
		t.SpentOn = c.input.SpentOn
		err := client.saveTimeEntry(t)

		if err == ContentError {
			fmt.Fprintf(c.errStream, "invalid content: %#v\n", t)
		} else if err != nil {
			return fmt.Errorf("failed to save time entry: %w", err)
		}

		bar.Add(1)
	}

	return nil
}

func (c *CMD) parseInput(file string) error {
	b, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to read a file: %w", err)
	}
	if err := yaml.Unmarshal(b, &c.input); err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}
	return nil
}
