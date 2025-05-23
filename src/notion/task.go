package notion

import "fmt"

type Task struct {
	Name        string
	ActionFlags []string
	Status      string
	Project     string
}

func (t *Task) String() string {
	return fmt.Sprintf("Name=%s, ActionFlags=%v, Status=%s, Project=%s", 
		t.Name, 
		t.ActionFlags, 
		t.Status, 
		t.Project,
	)
}
