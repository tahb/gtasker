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

func TestExecuteAction(t *testing.T) {
	mockClient := &MockTodoistClient{}

	executeAction(mockClient, "tasks")

	if !mockClient.GetTasksCalled {
		t.Errorf("Expected GetTasks to be called")
	}
}
