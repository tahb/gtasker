package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// TodoistAPIClient is a struct that represents a Todoist API client
type TodoistAPIClient struct {
	projectId string
	token     string
}

// TodoistClient is an interface that represents a Todoist client
type TodoistClient interface {
	GetTasks() ([]Task, error)
	CreateTask(project_id string, description string, content string) (Task, error)
}

// Due is a struct that represents a Todoist due date
type Due struct {
	Date        string `json:"date"`
	IsRecurring bool   `json:"is_recurring"`
	Datetime    string `json:"datetime"`
	String      string `json:"string"`
	Timezone    string `json:"timezone"`
}

// Task is a struct that represents a Todoist task
type Task struct {
	CreatorID    string   `json:"creator_id"`
	CreatedAt    string   `json:"created_at"`
	AssigneeID   string   `json:"assignee_id"`
	AssignerID   string   `json:"assigner_id"`
	CommentCount int      `json:"comment_count"`
	IsCompleted  bool     `json:"is_completed"`
	Content      string   `json:"content"`
	Description  string   `json:"description"`
	Due          Due      `json:"due"`
	Duration     *int     `json:"duration"`
	ID           string   `json:"id"`
	Labels       []string `json:"labels"`
	Order        int      `json:"order"`
	Priority     int      `json:"priority"`
	ProjectID    string   `json:"project_id"`
	SectionID    string   `json:"section_id"`
	ParentID     string   `json:"parent_id"`
	URL          string   `json:"url"`
}

// GetTasks is a method that returns a list of Todoist tasks
func (client *TodoistAPIClient) GetTasks() ([]Task, error) {
	url := "https://api.todoist.com/rest/v2/tasks?project_id=" + client.projectId

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []Task{}, err
	}

	req.Header.Set("Authorization", "Bearer "+client.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Task{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return []Task{}, err
	}

	var tasks []Task
	err = json.Unmarshal(body, &tasks)

	if err != nil {
		return []Task{}, err
	}

	return tasks, nil
}

func (client *TodoistAPIClient) CreateTask(project_id string, description string, content string) (Task, error) {
	return Task{}, nil
}
