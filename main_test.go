package main

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	main()

	expected := "action:tasks"
	got := buf.String()

	if !strings.Contains(got, expected) {
		t.Errorf("Expected output to contain: %q, got: %q", expected, got)
	}
}

type MockClient struct{}

func (m *MockClient) GetTasks() ([]Task, error) {
	return []Task{Task{ID: "1"}, Task{ID: "2"}, Task{ID: "3"}}, nil
}

func TestExecuteAction(t *testing.T) {
	mockClient := &MockClient{}

	var buf bytes.Buffer
	log.SetOutput(&buf)

	executeAction(mockClient, "tasks")

	expectedTasks := []Task{
		Task{ID: "1"},
		Task{ID: "2"},
		Task{ID: "3"},
	}
	expectedTaskIds := []string{}
	for _, task := range expectedTasks {
		expectedTaskIds = append(expectedTaskIds, task.ID)
	}

	got := buf.String()

	for _, expectedTaskId := range expectedTaskIds {
		if !strings.Contains(got, expectedTaskId) {
			t.Errorf("Expected output to contain: %q, got: %q", expectedTaskId, got)
		}
	}
}
