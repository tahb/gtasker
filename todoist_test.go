package main

import (
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func TestNewTodoistAPIClient(t *testing.T) {
	projectId := "your-project-id"
	token := "your-token"

	client := &TodoistAPIClient{
		projectId: projectId,
		token:     token,
	}

	if client.projectId != projectId {
		t.Errorf("Expected projectId to be %q, but got %q", projectId, client.projectId)
	}

	if client.token != token {
		t.Errorf("Expected token to be %q, but got %q", token, client.token)
	}
}

type MockTodoistClient struct {
	Task  Task
	Tasks []Task
	Err   error
	GetTasksCalled bool
}

func (m *MockTodoistClient) GetTasks() ([]Task, error) {
	m.GetTasksCalled = true
	return m.Tasks, m.Err
}

func (m *MockTodoistClient) CreateTask(project_id string, description string, content string) (Task, error) {
	return m.Task, m.Err
}

func TestGetTasks(t *testing.T) {
	mockClient := &MockTodoistClient{
		Tasks: []Task{newRandomTask(), newRandomTask(), newRandomTask()},
		Err:   nil,
	}

	tasks, err := mockClient.GetTasks()

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}

	if len(tasks) != len(mockClient.Tasks) {
		t.Errorf("Expected %d tasks, but got %d", len(mockClient.Tasks), len(tasks))
	}

	for i, task := range tasks {
		if task.CreatorID != mockClient.Tasks[i].CreatorID ||
			task.CreatedAt != mockClient.Tasks[i].CreatedAt ||
			task.AssigneeID != mockClient.Tasks[i].AssigneeID ||
			task.AssignerID != mockClient.Tasks[i].AssignerID ||
			task.CommentCount != mockClient.Tasks[i].CommentCount ||
			task.IsCompleted != mockClient.Tasks[i].IsCompleted ||
			task.Content != mockClient.Tasks[i].Content ||
			task.Description != mockClient.Tasks[i].Description ||
			!reflect.DeepEqual(task.Due, mockClient.Tasks[i].Due) ||
			!reflect.DeepEqual(task.Duration, mockClient.Tasks[i].Duration) ||
			task.ID != mockClient.Tasks[i].ID ||
			!reflect.DeepEqual(task.Labels, mockClient.Tasks[i].Labels) ||
			task.Order != mockClient.Tasks[i].Order ||
			task.Priority != mockClient.Tasks[i].Priority ||
			task.ProjectID != mockClient.Tasks[i].ProjectID ||
			task.SectionID != mockClient.Tasks[i].SectionID ||
			task.ParentID != mockClient.Tasks[i].ParentID ||
			task.URL != mockClient.Tasks[i].URL {
			t.Errorf("Expected task to be %v, but got %v", mockClient.Tasks[i], task)
		}
	}
}

func TestCreateTask(t *testing.T) {
	mockClient := &MockTodoistClient{
		Task: newRandomTask(),
		Err:  nil,
	}

	task, err := mockClient.CreateTask(
		"project_id",
		"description",
		"content",
	)

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}

	if task.ProjectID != mockClient.Task.ProjectID {
		t.Errorf("Expected project_id to be %q, but got %q", mockClient.Task.ProjectID, task.ProjectID)
	}
}

func newRandomTask() Task {
	duration := gofakeit.Number(1, 120)
	return Task{
		CreatorID:    gofakeit.UUID(),
		CreatedAt:    gofakeit.Date().Format("2006-01-02T15:04:05Z07:00"),
		AssigneeID:   gofakeit.UUID(),
		AssignerID:   gofakeit.UUID(),
		CommentCount: gofakeit.Number(1, 100),
		IsCompleted:  gofakeit.Bool(),
		Content:      gofakeit.Sentence(10),
		Description:  gofakeit.Paragraph(1, 2, 2, "\n"),
		Due: Due{
			Date:     gofakeit.Date().Format("2006-01-02"),
			Datetime: gofakeit.Date().Format("2006-01-02T15:04:05Z07:00"),
			String:   gofakeit.Sentence(5),
		},
		Duration:  &duration,
		ID:        gofakeit.UUID(),
		Labels:    []string{gofakeit.Word(), gofakeit.Word()},
		Order:     gofakeit.Number(1, 10),
		Priority:  gofakeit.Number(1, 4),
		ProjectID: gofakeit.UUID(),
		SectionID: gofakeit.UUID(),
		ParentID:  gofakeit.UUID(),
		URL:       gofakeit.URL(),
	}
}
