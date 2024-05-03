package main

import (
	"flag"
	"log"
)

func main() {
	projectId := flag.String("projectId", "", "a string")
	token := flag.String("token", "", "a string")
	action := flag.String("action", "tasks", "a string")

	flag.Parse()

	log.Print("projectId:", *projectId)
	log.Print("action:", *action)

	client := &TodoistAPIClient{projectId: *projectId, token: *token}

	executeAction(client, *action)
}

// executeAction is a method takes a Todoist client and an action and executes
// it against the TodoistAPI.
// - Currently only supports getting tasks
func executeAction(client TodoistClient, action string) {
	switch action {
	case "tasks":
		tasks, _ := client.GetTasks()

		for _, task := range tasks {
			log.Print(task)
		}
	}
}
